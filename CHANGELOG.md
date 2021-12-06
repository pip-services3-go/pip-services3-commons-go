
# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> Portable Abstractions and Patterns for Golang Changelog

## <a name="1.1.4"></a> 1.1.4 (2021-12-06)
### Bug Fixing 
* Fixed TypeConverter, ToTypeWithDefault method with nil default value crashed app
* Fixed AnyValueArray, GetAsMapWithDefault crashed app with stack overflow error

## <a name="1.1.3"></a> 1.1.3 (2021-10-22)
### Features 
* Added support access to object field by names in tags in reflect package

## <a name="1.1.2"></a> 1.1.2 (2021-09-20)

### Bug Fixing

* Fixed output of ObjectReader property names sortable lexicographically
## <a name="1.1.1"></a> 1.1.1 (2021-05-12)

### Bug Fixing

* Fix converting from uint64 to string and from string to uint64

### Features

* Added separated methods for conversion into signed and unsigned values 

## <a name="1.1.0"></a> 1.1.0 (2021-04-03)

### Features

* Moved to Go version 1.16
* Added JsonConverter.ToObject and ToObjectAs methods
* Added TokenizedPagingParams and TokenizedDataPage

## <a name="1.0.4"></a> 1.0.4 (2021-03-30)

### Bug fix

* Add json tags for PagingParams

## <a name="1.0.3"></a> 1.0.3 (2020-11-20)

### Features

* Add GetSchema method in command

## <a name="1.0.3"></a> 1.0.3 (2020-11-20)

### Features

* Add GetSchema method in command


## <a name="1.0.2"></a> 1.0.2 (2020-10-15)

### Bug Fixes
* Fix Long validation


## <a name="1.0.1"></a> 1.0.1 (2020-03-21)

### Bug Fixes
* Fixed doc typos


## <a name="1.0.0"></a> 1.0.0 (2020-01-28)

Initial public release

### Features
* **build** Component factories framework
* **commands** Command and Eventing patterns
* **config** Configuration framework
* **convert** Portable soft data converters
* **count** Performance counters components
* **data** Data value objects and random value generators
* **errors** Portable application errors
* **log** Logging components
* **random** Random data generators
* **refer** Component referencing framework
* **reflect** Portable reflection helpers
* **run** Execution framework
* **validate** Data validators
