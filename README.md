midievent [![GoDoc](https://godoc.org/github.com/xlab/midievent?status.svg)](https://godoc.org/github.com/xlab/midievent)
=========

Package `midievent` provides MIDI event type mappings for Go-lang. See docs:

* [MIDI_Code.pdf](/docs/MIDI_Code.pdf)
* [MIDI_Message_Table.pdf](/docs/MIDI_Message_Table.pdf)

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
