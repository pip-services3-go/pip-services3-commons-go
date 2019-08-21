# <img src="https://github.com/pip-services/pip-services/raw/master/design/Logo.png" alt="Pip.Services Logo" style="max-width:30%"> <br/> Portable Abstractions and Patterns for Node.js

This framework is part of the [Pip.Services](https://github.com/pip-services/pip-services) project.
It provides portable abstractions and patterns that can be used to implement non-trivial business logic in applications and services.

This framework's key difference is its portable implementation across a variety of different languages. 
It currently supports Java, .NET, Python, Node.js, and Golang. The code provides a reasonably thin abstraction layer 
over most fundamental functions and delivers symmetric implementation that can be quickly ported between different platforms.

The framework's functionality is decomposed into several packages:

- [**Commands**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/commands) - commanding and eventing patterns
- [**Config**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/config) - configuration framework
- [**Convert**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/convert) - soft value converters
- [**Data**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/data) - data patterns
- [**Errors**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/errors) - application errors
- [**Random**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/random) - random data generators
- [**Refer**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/refer) - locator (IoC) pattern
- [**Reflect**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/reflect) - reflection framework
- [**Run**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/run) - execution framework
- [**Validate**](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go/validate) - validation framework

Quick Links:

* [Downloads]
* [API Reference](https://godoc.org/github.com/pip-services3-go/pip-services3-commons-go)
* [Building and Testing]
* [Contributing]

## Acknowledgements

The library is created and maintained by **Sergey Seroukhov**.

The documentation is written by **Danyil Tretiakov**.
