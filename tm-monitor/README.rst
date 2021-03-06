Monitoring
==========

tm-monitor
----------

Tendermint blockchain monitoring tool; watches over one or more nodes, collecting and providing various statistics to the user: https://github.com/tendermint/tools/tree/master/tm-monitor

Quick Start
^^^^^^^^^^^

Docker
~~~~~~

Assuming your application is running in another container with the name ``app``:

::

    docker run -it --rm -v "/tmp:/tendermint" tendermint/tendermint init
    docker run -it --rm -v "/tmp:/tendermint" -p "46657:46657" --name=tm --link=app tendermint/tendermint node --proxy_app=tcp://app:46658

    docker run -it --rm -p "46670:46670" --link=tm tendermint/monitor tm:46657

If you don't have an application yet, but still want to try monitor out, use ``kvstore``:

::

    docker run -it --rm -v "/tmp:/tendermint" tendermint/tendermint init
    docker run -it --rm -v "/tmp:/tendermint" -p "46657:46657" --name=tm tendermint/tendermint node --proxy_app=kvstore

    docker run -it --rm -p "46670:46670" --link=tm tendermint/monitor tm:46657

Binaries
~~~~~~~~

::

    tm-monitor localhost:46657

Build from source
~~~~~~~~~~~~~~~~~

::

    make get_tools
    make get_vendor_deps
    make install

Usage
^^^^^

::

    tm-monitor [-v] [-no-ton] [-listen-addr="tcp://0.0.0.0:46670"] [endpoints]

    Examples:
            # monitor single instance
            tm-monitor localhost:46657

            # monitor a few instances by providing comma-separated list of RPC endpoints
            tm-monitor host1:46657,host2:46657
    Flags:
      -listen-addr string
            HTTP and Websocket server listen address (default "tcp://0.0.0.0:46670")
      -no-ton
            Do not show ton (table of nodes)
      -v    verbose logging

RPC UI
^^^^^^

Run ``tm-monitor`` and visit http://localhost:46670
You should see the list of the available RPC endpoints:

::

    http://localhost:46670/status
    http://localhost:46670/status/network
    http://localhost:46670/monitor?endpoint=_
    http://localhost:46670/status/node?name=_
    http://localhost:46670/unmonitor?endpoint=_

The API is available as GET requests with URI encoded parameters, or as JSONRPC
POST requests. The JSONRPC methods are also exposed over websocket.

Development
^^^^^^^^^^^

::

    make get_tools
    make get_vendor_deps
    make test
