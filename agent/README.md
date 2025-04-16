# agent
prometheus agent 是一个独立的程序，用于将prometheus中的告警规则实时同步到[pulse](https://gitee.com/florenbai_1_floren_bai/pulse)中。如果在prometheus中配置告警规则，需要将告警规则实时同步到pulse中，需要使用该agent。

## 如何部署

在agent/config/config.yaml中配置:

- nacos : 集成了nacos作为服务发现，请详细配置这些信息
- prometheus : 集成了prometheus作为告警规则的来源，请详细配置这些信息。sourceSign : 用于标识prometheus的来源，需要与pulse中告警源的配置一致。

```bash
go build -o agent main.go
./agent
```