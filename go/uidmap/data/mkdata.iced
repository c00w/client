{Gets} = require('iced-utils')
{make_esc} = require 'iced-error'

gojoin = (v) -> v.join(", ") + ","

uid_to_byte_list = (uid) -> gojoin ("0x" + uid[i...(i+2)].toString(16) for i in [0...uid.length] by 2)

class Runner

  constructor : () ->
    map = []

  read : ({input}, cb) ->
    esc = make_esc cb, "read"
    gets = (new Gets input).run()
    map = []
    loop
      await gets.gets esc defer line
      break unless line?
      [uid,username] = line.split /\s+/
      map.push [ uid, username ]
    cb null, map

  output : ({map}, cb) ->
    out = []
    out.push "package uidmap"
    out.push ''
    @output_uid_list { map, out }
    out.push ''
    @output_username_block { map, out }
    out.push ''
    @output_lengths { map, out }
    out.join "\n"

  output_username_block : ({map, out}) ->
    v = "var usernames = " + '"' + (username for [_,username] in map).join('') + '"'
    out.push v

  output_lengths : ({map, out}) ->
    out.push "var lengths = []uint8{"
    offset = 0
    for i in [0...map.length] by 16
      v = for [_,username] in map[i...(i+16)]
        username.length.toString 10
      out.push "\t" + gojoin v
    out.push "}"

  output_uid_list : ({map, out}) ->
    out.push "var uids = []byte{"
    for [uid] in map
      out.push "\t" + uid_to_byte_list(uid)
    out.push "}"

  run : ({input}, cb) ->
    esc = make_esc cb, "run"
    await @read { input }, esc defer map
    out = @output { map }
    console.log out
    cb null

r = new Runner
await r.run { input : process.stdin }, defer err
rc = 0
if err?
  console.err err.toString()
  rc = -2
process.exit rc
