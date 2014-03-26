waiter
------

Simple process to wait for SIGINT/SIGTERM before exiting.

    $ ./waiter -h
    Usage of ./waiter:
      -int=1: how long to wait before exiting after receiving SIGINT
      -term=1: how long to wait before exiting after receiving SIGTERM

An infinite loop until you send it a signal.

Install
-------

    $ go get github.com/stengaard/waiter


[![Gobuild Download](http://gobuild.io/badge/github.com/stengaard/waiter/download.png)](http://gobuild.io/github.com/stengaard/waiter)
