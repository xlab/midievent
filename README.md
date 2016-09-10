midievent
=========

Package `midievent` provides MIDI event type mappings for Go-lang.

## Install

```
$ go get github.com/xlab/midievent
```

## Use

```go
for ev := range midiIn.Source() {
    msg := portmidi.Message(ev.Message)
    if midievent.IsNoteOn(midievent.Event(msg.Status())) {
        n := int(msg.Data1())
        log.Printf("note %d (%.3fHz)", n, noteToFreq(n))
    }
}
```

Example: [vocoder](https://github.com/xlab/portmidi/tree/master/example/vocoder).

## License

CC0 (public domain).
