package main

import (
    "fmt"
    "log"
    "time"

    hd44780 "github.com/adrianh-za/go-hd44780-rpi"
    "github.com/d2r2/go-i2c"
    logger "github.com/d2r2/go-logger"
)

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    //Stop the I2C module from spamming the console
    logger.ChangePackageLogLevel("i2c", logger.InfoLevel)

    //Init I2C
    i2c, err := i2c.NewI2C(0x27, 1)
    checkError(err)
    defer i2c.Close()

    //Init the display
    lcd, err := hd44780.NewLcd(i2c, hd44780.LCD_20x4)
    lcd.Startup()
    lcd.SetupExit(true) //Setup CTRL-C to quit gracefully

    //Display the time
    for {
        lcd.Home()
        t := time.Now()
        seconds := t.Second()

        lcd.SetPosition(0, 0)
        fmt.Fprint(lcd, t.Format("Monday Jan 2"))
        lcd.SetPosition(1, 1)
        fmt.Fprint(lcd, t.Format("15:04:05 2006"))
        lcd.SetPosition(2, 0)
        fmt.Fprint(lcd, t.Format("Monday Jan 2"))
        lcd.SetPosition(3, 1)
        fmt.Fprint(lcd, t.Format("15:04:05 2006"))

        //Small sleeps until we need to update due to second change.
        for {
            if seconds != time.Now().Second() {
                break
            }

            time.Sleep(200 * time.Millisecond)
        }
    }
}
