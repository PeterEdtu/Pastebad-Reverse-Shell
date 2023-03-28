# Pastebad Reverse Shell
Simple Cross-platform Reverse Shell using Pastebin website as C2.

This project is meant for MalDev training and PoC only, not for spicy things. I'm not responsible for what you're doing with this code.

Anyways, don't forget to use a TOR engine when massively testing this malware PoC e.g. [anonsurf](https://github.com/Und3rf10w/kali-anonsurf) or [nipe](https://github.com/htrgouvea/nipe).

## How it works
The malware simply requests Pastebin's text from an ID which is incremented each time to get the next command. For now, no exclusions are made, so everyone can create a new paste text, and thus execute the command automatically on the infected device. I chose [pastebin.fr](http://pastebin.fr/) because it's perfectly working with the TOR nodes (no ip blacklist) and doesn't seem to be flagged by EDR nor "enterprise" VPN/Proxy yet.

The commands are Base64 encoded and executed if the decode works, and shall not crash the program if an error occurs (need more tests on cases/operating systems). All tests for now are made on my Kali Linux VM.

I made this project for fun, I don't know if this will be any more serious than that, but I'm looking forward to learn more about MalDev to give constructive contribution to the Red / Blue Team community :)

## Actual features
- Full implementation in Golang working on Linux and Windows
- Fixed-delay Reverse Shell beacon to [pastebin.fr](http://pastebin.fr/) using encoded commands (Base64)
- Simple strings obfuscation for the URL

## Next steps
- Implementation in Nim
- Random delay for beaconing
- Command encryption?
- Code optimisation and build compression (for now it's quite a chunky boi)
- Payload integration
- More MalDev stuff which I don't know yet
