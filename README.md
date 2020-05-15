# TryDo

A Go Try-Do library.

It's simple and easy to use.

## Examples

```
// try 3 times
var i int32 = 0
do := func() error {
    c := atomic.AddInt32(&i, 1)
    if c < 3 {
        return errors.New("err")
    }
    return nil
}

err := trydo.TryTimes(do, 3)

// try 3 times with intervals
err := trydo.TryWithIntervals(do, time.Second, 2*time.Second, 5*time.Second)
```