
## CLI конвертирует одну валюту в другую используя web API https://exchangeratesapi.io
Build:

`go build -o fiatconv cmd/fiatconv/main.go`

Параметры командной строки: <amount:float> <src_symbol:string> <dst_symbol:string>

`./fiatconv 123.45 USD RUB`