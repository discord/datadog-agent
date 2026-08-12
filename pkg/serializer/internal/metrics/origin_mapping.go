// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package metrics

import (
	"github.com/DataDog/datadog-agent/pkg/metrics"
)

func metricSourceToOriginProduct(ms metrics.MetricSource) int32 {
	const serieMetadataOriginOriginProductServerlessType = 1
	const serieMetadataOriginOriginProductAgentType = 10
	const serieMetadataOriginOriginProductDatadogExporterType = 19
	const serieMetadataOriginOriginProductGPU = 38 // ref: https://github.com/DataDog/dd-source/blob/276882b71d84785ec89c31973046ab66d5a01807/domains/metrics/shared/libs/proto/origin/origin.proto#L277
	if ms >= metrics.MetricSourceOpenTelemetryCollectorUnknown && ms <= metrics.MetricSourceOpenTelemetryCollectorCouchdbReceiver {
		return serieMetadataOriginOriginProductDatadogExporterType
	}
	if ms == metrics.MetricSourceGPU {
		return serieMetadataOriginOriginProductGPU
	}
	switch ms {
	case metrics.MetricSourceServerless,
		metrics.MetricSourceAzureContainerAppCustom,
		metrics.MetricSourceAzureContainerAppEnhanced,
		metrics.MetricSourceAzureContainerAppRuntime,
		metrics.MetricSourceAzureAppServiceCustom,
		metrics.MetricSourceAzureAppServiceEnhanced,
		metrics.MetricSourceAzureAppServiceRuntime,
		metrics.MetricSourceGoogleCloudRunCustom,
		metrics.MetricSourceGoogleCloudRunEnhanced,
		metrics.MetricSourceGoogleCloudRunRuntime:
		return serieMetadataOriginOriginProductServerlessType
	}
	return serieMetadataOriginOriginProductAgentType
}

func metricSourceToOriginCategory(ms metrics.MetricSource) int32 {
	// Extreme hack to handle the passthrough of metadata.origin.service from the datadog receiver
	// (the highest bit of the source is set when an origin service was detected in the receiver)
	// (comp/otelcol/otlp/components/exporter/serializerexporter/consumer.go)
	if ms&(1<<15) != 0 {
		// the origin service is simply the lower 15 bits of the source
		// so we need to convert it back to a MetricSource for this function to work correctly
		ms = originServiceToMetricSource(int32(uint16(ms) & ^(uint16(1) << 15)))
	}
	// These constants map to specific fields in the 'OriginSubproduct' enum in origin.proto
	switch ms {
	case metrics.MetricSourceUnknown:
		return 0
	case metrics.MetricSourceDogstatsd:
		return 10
	case metrics.MetricSourceJmxCustom,
		metrics.MetricSourceActivemq,
		metrics.MetricSourceAnyscale,
		metrics.MetricSourceAppgateSDP,
		metrics.MetricSourceCassandra,
		metrics.MetricSourceConfluentPlatform,
		metrics.MetricSourceHazelcast,
		metrics.MetricSourceHive,
		metrics.MetricSourceHivemq,
		metrics.MetricSourceHudi,
		metrics.MetricSourceIgnite,
		metrics.MetricSourceJbossWildfly,
		metrics.MetricSourceKafka,
		metrics.MetricSourcePresto,
		metrics.MetricSourceSolr,
		metrics.MetricSourceSonarqube,
		metrics.MetricSourceTomcat,
		metrics.MetricSourceWeblogic,
		// Core Checks
		metrics.MetricSourceInternal,
		metrics.MetricSourceContainer,
		metrics.MetricSourceContainerd,
		metrics.MetricSourceCri,
		metrics.MetricSourceDocker,
		metrics.MetricSourceNTP,
		metrics.MetricSourceSystemd,
		metrics.MetricSourceHelm,
		metrics.MetricSourceKubeflow,
		metrics.MetricSourceKubernetesAPIServer,
		metrics.MetricSourceKubernetesStateCore,
		metrics.MetricSourceOrchestrator,
		metrics.MetricSourceWinproc,
		metrics.MetricSourceFileHandle,
		metrics.MetricSourceWinkmem,
		metrics.MetricSourceIo,
		metrics.MetricSourceUptime,
		metrics.MetricSourceSbom,
		metrics.MetricSourceMemory,
		metrics.MetricSourceTCPQueueLength,
		metrics.MetricSourceOomKill,
		metrics.MetricSourceContainerLifecycle,
		metrics.MetricSourceJetson,
		metrics.MetricSourceContainerImage,
		metrics.MetricSourceCPU,
		metrics.MetricSourceLoad,
		metrics.MetricSourceDisk,
		metrics.MetricSourceNetwork,
		metrics.MetricSourceSnmp,
		metrics.MetricSourceCiscoSdwan,
		metrics.MetricSourceVersa,
		metrics.MetricSourceWlan,
		metrics.MetricSourceWindowsCertificateStore,
		// Plugins and non-checks
		metrics.MetricSourceCloudFoundry,
		metrics.MetricSourceJenkins,
		// Python Checks
		metrics.MetricSourceZenohRouter,
		metrics.MetricSourceZabbix,
		metrics.MetricSourceWayfinder,
		metrics.MetricSourceVespa,
		metrics.MetricSourceUpsc,
		metrics.MetricSourceUpboundUxp,
		metrics.MetricSourceUnifiConsole,
		metrics.MetricSourceUnbound,
		metrics.MetricSourceTraefik,
		metrics.MetricSourceTidb,
		metrics.MetricSourceSyncthing,
		metrics.MetricSourceStorm,
		metrics.MetricSourceStardog,
		metrics.MetricSourceSpeedtest,
		metrics.MetricSourceSortdb,
		metrics.MetricSourceSonarr,
		metrics.MetricSourceSnmpwalk,
		metrics.MetricSourceSendmail,
		metrics.MetricSourceScaphandre,
		metrics.MetricSourceScalr,
		metrics.MetricSourceRiakRepl,
		metrics.MetricSourceRedpanda,
		metrics.MetricSourceRedisenterprise,
		metrics.MetricSourceRedisSentinel,
		metrics.MetricSourceRebootRequired,
		metrics.MetricSourceRadarr,
		metrics.MetricSourcePurefb,
		metrics.MetricSourcePurefa,
		metrics.MetricSourcePuma,
		metrics.MetricSourcePortworx,
		metrics.MetricSourcePing,
		metrics.MetricSourcePihole,
		metrics.MetricSourcePhpOpcache,
		metrics.MetricSourcePhpApcu,
		metrics.MetricSourceOpenPolicyAgent,
		metrics.MetricSourceOctopusDeploy,
		metrics.MetricSourceOctoprint,
		metrics.MetricSourceNvml,
		metrics.MetricSourceNs1,
		metrics.MetricSourceNnSdwan,
		metrics.MetricSourceNextcloud,
		metrics.MetricSourceNeutrona,
		metrics.MetricSourceNeo4j,
		metrics.MetricSourceMergify,
		metrics.MetricSourceLogstash,
		metrics.MetricSourceLighthouse,
		metrics.MetricSourceKernelcare,
		metrics.MetricSourceKepler,
		metrics.MetricSourceJfrogPlatformSelfHosted,
		metrics.MetricSourceHikaricp,
		metrics.MetricSourceGrpcCheck,
		metrics.MetricSourceGoPprofScraper,
		metrics.MetricSourceGnatsdStreaming,
		metrics.MetricSourceGnatsd,
		metrics.MetricSourceGitea,
		metrics.MetricSourceGatekeeper,
		metrics.MetricSourceFlyIo,
		metrics.MetricSourceFluentbit,
		metrics.MetricSourceFilemage,
		metrics.MetricSourceFilebeat,
		metrics.MetricSourceFiddler,
		metrics.MetricSourceExim,
		metrics.MetricSourceEventstore,
		metrics.MetricSourceEmqx,
		metrics.MetricSourceCyral,
		metrics.MetricSourceCybersixgillActionableAlerts,
		metrics.MetricSourceCloudsmith,
		metrics.MetricSourceCloudnatix,
		metrics.MetricSourceCfssl,
		metrics.MetricSourceBind9,
		metrics.MetricSourceAwsPricing,
		metrics.MetricSourceAqua,
		metrics.MetricSourceKubernetesClusterAutoscaler,
		metrics.MetricSourceTraefikMesh,
		metrics.MetricSourceWeaviate,
		metrics.MetricSourceTorchserve,
		metrics.MetricSourceTemporal,
		metrics.MetricSourceTeleport,
		metrics.MetricSourceTekton,
		metrics.MetricSourceStrimzi,
		metrics.MetricSourceRay,
		metrics.MetricSourceNvidiaTriton,
		metrics.MetricSourceKarpenter,
		metrics.MetricSourceKubeVirtAPI,
		metrics.MetricSourceKubeVirtController,
		metrics.MetricSourceKubeVirtHandler,
		metrics.MetricSourceFluxcd,
		metrics.MetricSourceEsxi,
		metrics.MetricSourceDcgm,
		metrics.MetricSourceDatadogClusterAgent,
		metrics.MetricSourceCloudera,
		metrics.MetricSourceArgoWorkflows,
		metrics.MetricSourceArgoRollouts,
		metrics.MetricSourceActiveDirectory,
		metrics.MetricSourceActivemqXML,
		metrics.MetricSourceAerospike,
		metrics.MetricSourceAirflow,
		metrics.MetricSourceAmazonMsk,
		metrics.MetricSourceAmbari,
		metrics.MetricSourceApache,
		metrics.MetricSourceArangodb,
		metrics.MetricSourceArgocd,
		metrics.MetricSourceAspdotnet,
		metrics.MetricSourceAviVantage,
		metrics.MetricSourceAzureIotEdge,
		metrics.MetricSourceBoundary,
		metrics.MetricSourceBtrfs,
		metrics.MetricSourceCacti,
		metrics.MetricSourceCalico,
		metrics.MetricSourceCassandraNodetool,
		metrics.MetricSourceCeph,
		metrics.MetricSourceCertManager,
		metrics.MetricSourceCilium,
		metrics.MetricSourceCitrixHypervisor,
		metrics.MetricSourceClickhouse,
		metrics.MetricSourceCloudFoundryAPI,
		metrics.MetricSourceCockroachdb,
		metrics.MetricSourceConsul,
		metrics.MetricSourceControlM,
		metrics.MetricSourceCoredns,
		metrics.MetricSourceCouch,
		metrics.MetricSourceCouchbase,
		metrics.MetricSourceCrio,
		metrics.MetricSourceDirectory,
		metrics.MetricSourceDNSCheck,
		metrics.MetricSourceDotnetclr,
		metrics.MetricSourceDruid,
		metrics.MetricSourceEcsFargate,
		metrics.MetricSourceEksFargate,
		metrics.MetricSourceElastic,
		metrics.MetricSourceEnvoy,
		metrics.MetricSourceEtcd,
		metrics.MetricSourceExchangeServer,
		metrics.MetricSourceExternalDNS,
		metrics.MetricSourceFluentd,
		metrics.MetricSourceFoundationdb,
		metrics.MetricSourceGearmand,
		metrics.MetricSourceGitlab,
		metrics.MetricSourceGitlabRunner,
		metrics.MetricSourceGlusterfs,
		metrics.MetricSourceGoExpvar,
		metrics.MetricSourceGunicorn,
		metrics.MetricSourceHaproxy,
		metrics.MetricSourceHarbor,
		metrics.MetricSourceHdfsDatanode,
		metrics.MetricSourceHdfsNamenode,
		metrics.MetricSourceHTTPCheck,
		metrics.MetricSourceHyperv,
		metrics.MetricSourceIbmAce,
		metrics.MetricSourceIbmDb2,
		metrics.MetricSourceIbmI,
		metrics.MetricSourceIbmMq,
		metrics.MetricSourceIbmWas,
		metrics.MetricSourceIis,
		metrics.MetricSourceImpala,
		metrics.MetricSourceIstio,
		metrics.MetricSourceKafkaConsumer,
		metrics.MetricSourceKong,
		metrics.MetricSourceKubeAPIserverMetrics,
		metrics.MetricSourceKubeControllerManager,
		metrics.MetricSourceKubeDNS,
		metrics.MetricSourceKubeMetricsServer,
		metrics.MetricSourceKubeProxy,
		metrics.MetricSourceKubeScheduler,
		metrics.MetricSourceKubelet,
		metrics.MetricSourceKubernetesState,
		metrics.MetricSourceKyototycoon,
		metrics.MetricSourceKyverno,
		metrics.MetricSourceLighttpd,
		metrics.MetricSourceLinkerd,
		metrics.MetricSourceLinuxProcExtras,
		metrics.MetricSourceMapr,
		metrics.MetricSourceMapreduce,
		metrics.MetricSourceMarathon,
		metrics.MetricSourceMarklogic,
		metrics.MetricSourceMcache,
		metrics.MetricSourceMesosMaster,
		metrics.MetricSourceMesosSlave,
		metrics.MetricSourceMongo,
		metrics.MetricSourceMysql,
		metrics.MetricSourceN8N,
		metrics.MetricSourceNutanix,
		metrics.MetricSourcePaloAltoPanorama,
		metrics.MetricSourcePrefect,
		metrics.MetricSourceNagios,
		metrics.MetricSourceNfsstat,
		metrics.MetricSourceNginx,
		metrics.MetricSourceNginxIngressController,
		metrics.MetricSourceOpenldap,
		metrics.MetricSourceOpenmetrics,
		metrics.MetricSourceOpenstack,
		metrics.MetricSourceOpenstackController,
		metrics.MetricSourceOracle,
		metrics.MetricSourcePdhCheck,
		metrics.MetricSourcePgbouncer,
		metrics.MetricSourcePhpFpm,
		metrics.MetricSourcePostfix,
		metrics.MetricSourcePostgres,
		metrics.MetricSourcePowerdnsRecursor,
		metrics.MetricSourceProcess,
		metrics.MetricSourcePrometheus,
		metrics.MetricSourceProxysql,
		metrics.MetricSourcePulsar,
		metrics.MetricSourceRabbitmq,
		metrics.MetricSourceRedisdb,
		metrics.MetricSourceRethinkdb,
		metrics.MetricSourceRiak,
		metrics.MetricSourceRiakcs,
		metrics.MetricSourceSapHana,
		metrics.MetricSourceScylla,
		metrics.MetricSourceSilk,
		metrics.MetricSourceSinglestore,
		metrics.MetricSourceSlurm,
		metrics.MetricSourceSnowflake,
		metrics.MetricSourceSpark,
		metrics.MetricSourceSqlserver,
		metrics.MetricSourceSquid,
		metrics.MetricSourceSSHCheck,
		metrics.MetricSourceStatsd,
		metrics.MetricSourceSupervisord,
		metrics.MetricSourceSystemCore,
		metrics.MetricSourceSystemSwap,
		metrics.MetricSourceTCPCheck,
		metrics.MetricSourceTeamcity,
		metrics.MetricSourceTeradata,
		metrics.MetricSourceTibcoEMS,
		metrics.MetricSourceTLS,
		metrics.MetricSourceTokumx,
		metrics.MetricSourceTrafficServer,
		metrics.MetricSourceTwemproxy,
		metrics.MetricSourceTwistlock,
		metrics.MetricSourceVarnish,
		metrics.MetricSourceVault,
		metrics.MetricSourceVertica,
		metrics.MetricSourceVelero,
		metrics.MetricSourceVllm,
		metrics.MetricSourceVoltdb,
		metrics.MetricSourceVsphere,
		metrics.MetricSourceWin32EventLog,
		metrics.MetricSourceWindowsPerformanceCounters,
		metrics.MetricSourceWindowsService,
		metrics.MetricSourceWmiCheck,
		metrics.MetricSourceYarn,
		metrics.MetricSourceZk,
		metrics.MetricSourceAwsNeuron,
		metrics.MetricSourceNvidiaNim,
		metrics.MetricSourceQuarkus,
		metrics.MetricSourceMilvus,
		metrics.MetricSourceCelery,
		metrics.MetricSourceInfiniband,
		metrics.MetricSourceAnecdote,
		metrics.MetricSourceSonatypeNexus,
		metrics.MetricSourceSilverstripeCMS,
		metrics.MetricSourceAltairPBSPro,
		metrics.MetricSourceFalco,
		metrics.MetricSourceKrakenD,
		metrics.MetricSourceKuma,
		metrics.MetricSourceLiteLLM,
		metrics.MetricSourceLustre,
		metrics.MetricSourceProxmox,
		metrics.MetricSourceSupabase,
		metrics.MetricSourceKeda,
		metrics.MetricSourceDuckdb,
		metrics.MetricSourceResilience4j,
		metrics.MetricSourceBentoMl,
		metrics.MetricSourceHuggingFaceTgi,
		metrics.MetricSourceIbmSpectrumLsf,
		metrics.MetricSourceDatadogOperator,
		metrics.MetricSourceBattery,
		metrics.MetricSourcePinot,
		metrics.MetricSourceDellPowerFlex,
		metrics.MetricSourceHPEArubaEdgeConnect,
		metrics.MetricSourceNiFi:
		return 11 // integrationMetrics
	case metrics.MetricSourceGPU:
		return 72 // ref: https://github.com/DataDog/dd-source/blob/276882b71d84785ec89c31973046ab66d5a01807/domains/metrics/shared/libs/proto/origin/origin.proto#L427
	case metrics.MetricSourceAzureAppServiceCustom,
		metrics.MetricSourceAzureAppServiceEnhanced,
		metrics.MetricSourceAzureAppServiceRuntime:
		return 35
	case metrics.MetricSourceGoogleCloudRunCustom,
		metrics.MetricSourceGoogleCloudRunEnhanced,
		metrics.MetricSourceGoogleCloudRunRuntime:
		return 36
	case metrics.MetricSourceAzureContainerAppCustom,
		metrics.MetricSourceAzureContainerAppEnhanced,
		metrics.MetricSourceAzureContainerAppRuntime:
		return 37
	default:
		return 0
	}
}

func metricSourceToOriginService(ms metrics.MetricSource) int32 {
	// Extreme hack to handle the passthrough of metadata.origin.service from the datadog receiver
	// (the highest bit of the source is set when an origin service was detected in the receiver)
	// (comp/otelcol/otlp/components/exporter/serializerexporter/consumer.go)
	if ms&(1<<15) != 0 {
		// the origin service is simply the lower 15 bits of the source
		return int32(uint16(ms) & ^(uint16(1) << 15))
	}
	// These constants map to specific fields in the 'OriginService' enum in origin.proto
	switch ms {
	case metrics.MetricSourceDogstatsd:
		return 0
	case metrics.MetricSourceJmxCustom:
		return 9
	case metrics.MetricSourceUnknown:
		return 0
	case metrics.MetricSourceActiveDirectory:
		return 10
	case metrics.MetricSourceActivemqXML:
		return 11
	case metrics.MetricSourceActivemq:
		return 12
	case metrics.MetricSourceAerospike:
		return 13
	case metrics.MetricSourceAirflow:
		return 14
	case metrics.MetricSourceAmazonMsk:
		return 15
	case metrics.MetricSourceAmbari:
		return 16
	case metrics.MetricSourceApache:
		return 17
	case metrics.MetricSourceArangodb:
		return 18
	case metrics.MetricSourceArgocd:
		return 19
	case metrics.MetricSourceAspdotnet:
		return 20
	case metrics.MetricSourceAviVantage:
		return 21
	case metrics.MetricSourceAzureIotEdge:
		return 22
	case metrics.MetricSourceBoundary:
		return 23
	case metrics.MetricSourceBtrfs:
		return 24
	case metrics.MetricSourceCacti:
		return 25
	case metrics.MetricSourceCalico:
		return 26
	case metrics.MetricSourceCassandraNodetool:
		return 27
	case metrics.MetricSourceCassandra:
		return 28
	case metrics.MetricSourceCeph:
		return 29
	case metrics.MetricSourceCertManager:
		return 30
	case metrics.MetricSourceCilium:
		return 34
	case metrics.MetricSourceCitrixHypervisor:
		return 36
	case metrics.MetricSourceClickhouse:
		return 37
	case metrics.MetricSourceCloudFoundry:
		return 440
	case metrics.MetricSourceCloudFoundryAPI:
		return 38
	case metrics.MetricSourceCockroachdb:
		return 39
	case metrics.MetricSourceConfluentPlatform:
		return 40
	case metrics.MetricSourceConsul:
		return 41
	case metrics.MetricSourceCoredns:
		return 42
	case metrics.MetricSourceCouch:
		return 43
	case metrics.MetricSourceCouchbase:
		return 44
	case metrics.MetricSourceCrio:
		return 45
	case metrics.MetricSourceDirectory:
		return 47
	case metrics.MetricSourceDisk:
		return 48
	case metrics.MetricSourceDNSCheck:
		return 49
	case metrics.MetricSourceDotnetclr:
		return 50
	case metrics.MetricSourceDruid:
		return 51
	case metrics.MetricSourceEcsFargate:
		return 52
	case metrics.MetricSourceEksFargate:
		return 53
	case metrics.MetricSourceElastic:
		return 54
	case metrics.MetricSourceEnvoy:
		return 55
	case metrics.MetricSourceEtcd:
		return 56
	case metrics.MetricSourceExchangeServer:
		return 57
	case metrics.MetricSourceExternalDNS:
		return 58
	case metrics.MetricSourceFluentd:
		return 60
	case metrics.MetricSourceFlyIo:
		return 430
	case metrics.MetricSourceFoundationdb:
		return 61
	case metrics.MetricSourceGearmand:
		return 62
	case metrics.MetricSourceGitlabRunner:
		return 63
	case metrics.MetricSourceGitlab:
		return 64
	case metrics.MetricSourceGlusterfs:
		return 65
	case metrics.MetricSourceGoExpvar:
		return 66
	case metrics.MetricSourceGPU:
		return 466
	case metrics.MetricSourceGunicorn:
		return 67
	case metrics.MetricSourceHaproxy:
		return 68
	case metrics.MetricSourceHarbor:
		return 69
	case metrics.MetricSourceHazelcast:
		return 70
	case metrics.MetricSourceHdfsDatanode:
		return 71
	case metrics.MetricSourceHdfsNamenode:
		return 72
	case metrics.MetricSourceHive:
		return 73
	case metrics.MetricSourceHivemq:
		return 74
	case metrics.MetricSourceHTTPCheck:
		return 75
	case metrics.MetricSourceHudi:
		return 76
	case metrics.MetricSourceHyperv:
		return 77
	case metrics.MetricSourceIbmAce:
		return 78
	case metrics.MetricSourceIbmDb2:
		return 79
	case metrics.MetricSourceIbmI:
		return 80
	case metrics.MetricSourceIbmMq:
		return 81
	case metrics.MetricSourceIbmWas:
		return 82
	case metrics.MetricSourceIgnite:
		return 83
	case metrics.MetricSourceIis:
		return 84
	case metrics.MetricSourceImpala:
		return 85
	case metrics.MetricSourceIstio:
		return 86
	case metrics.MetricSourceJbossWildfly:
		return 87
	case metrics.MetricSourceJenkins:
		return 436
	case metrics.MetricSourceKafkaConsumer:
		return 89
	case metrics.MetricSourceKafka:
		return 90
	case metrics.MetricSourceKepler:
		return 431
	case metrics.MetricSourceKong:
		return 91
	case metrics.MetricSourceKubeAPIserverMetrics:
		return 92
	case metrics.MetricSourceKubeControllerManager:
		return 93
	case metrics.MetricSourceKubeDNS:
		return 94
	case metrics.MetricSourceKubeMetricsServer:
		return 95
	case metrics.MetricSourceKubeProxy:
		return 96
	case metrics.MetricSourceKubeScheduler:
		return 97
	case metrics.MetricSourceKubelet:
		return 98
	case metrics.MetricSourceKubernetesState:
		return 99
	case metrics.MetricSourceKubeVirtAPI:
		return 437
	case metrics.MetricSourceKubeVirtController:
		return 438
	case metrics.MetricSourceKubeVirtHandler:
		return 439
	case metrics.MetricSourceKyototycoon:
		return 100
	case metrics.MetricSourceLighttpd:
		return 101
	case metrics.MetricSourceLinkerd:
		return 102
	case metrics.MetricSourceLinuxProcExtras:
		return 103
	case metrics.MetricSourceMapr:
		return 104
	case metrics.MetricSourceMapreduce:
		return 105
	case metrics.MetricSourceMarathon:
		return 106
	case metrics.MetricSourceMarklogic:
		return 107
	case metrics.MetricSourceMcache:
		return 108
	case metrics.MetricSourceMesosMaster:
		return 109
	case metrics.MetricSourceMesosSlave:
		return 110
	case metrics.MetricSourceMongo:
		return 111
	case metrics.MetricSourceMysql:
		return 112
	case metrics.MetricSourceNagios:
		return 113
	case metrics.MetricSourceNetwork:
		return 114
	case metrics.MetricSourceNfsstat:
		return 115
	case metrics.MetricSourceNginxIngressController:
		return 116
	case metrics.MetricSourceNginx:
		return 117
	case metrics.MetricSourceOctopusDeploy:
		return 432
	case metrics.MetricSourceOpenldap:
		return 118
	case metrics.MetricSourceOpenmetrics:
		return 119
	case metrics.MetricSourceOpenstackController:
		return 120
	case metrics.MetricSourceOpenstack:
		return 121
	case metrics.MetricSourceOracle:
		return 122
	case metrics.MetricSourcePdhCheck:
		return 124
	case metrics.MetricSourcePgbouncer:
		return 125
	case metrics.MetricSourcePhpFpm:
		return 126
	case metrics.MetricSourcePostfix:
		return 127
	case metrics.MetricSourcePostgres:
		return 128
	case metrics.MetricSourcePowerdnsRecursor:
		return 129
	case metrics.MetricSourcePresto:
		return 130
	case metrics.MetricSourceProcess:
		return 131
	case metrics.MetricSourcePrometheus:
		return 132
	case metrics.MetricSourceProxysql:
		return 133
	case metrics.MetricSourcePulsar:
		return 134
	case metrics.MetricSourceRabbitmq:
		return 135
	case metrics.MetricSourceRedisdb:
		return 136
	case metrics.MetricSourceRethinkdb:
		return 137
	case metrics.MetricSourceRiak:
		return 138
	case metrics.MetricSourceRiakcs:
		return 139
	case metrics.MetricSourceSapHana:
		return 140
	case metrics.MetricSourceScaphandre:
		return 433
	case metrics.MetricSourceScylla:
		return 141
	case metrics.MetricSourceSilk:
		return 143
	case metrics.MetricSourceSinglestore:
		return 144
	case metrics.MetricSourceSnmp:
		return 145
	case metrics.MetricSourceSnowflake:
		return 146
	case metrics.MetricSourceSolr:
		return 147
	case metrics.MetricSourceSonarqube:
		return 148
	case metrics.MetricSourceSpark:
		return 149
	case metrics.MetricSourceSqlserver:
		return 150
	case metrics.MetricSourceSquid:
		return 151
	case metrics.MetricSourceSSHCheck:
		return 152
	case metrics.MetricSourceStatsd:
		return 153
	case metrics.MetricSourceSupervisord:
		return 154
	case metrics.MetricSourceSystemCore:
		return 155
	case metrics.MetricSourceSystemSwap:
		return 156
	case metrics.MetricSourceTCPCheck:
		return 157
	case metrics.MetricSourceTeamcity:
		return 158
	case metrics.MetricSourceTeradata:
		return 160
	case metrics.MetricSourceTLS:
		return 161
	case metrics.MetricSourceTokumx:
		return 162
	case metrics.MetricSourceTomcat:
		return 163
	case metrics.MetricSourceTrafficServer:
		return 164
	case metrics.MetricSourceTwemproxy:
		return 165
	case metrics.MetricSourceTwistlock:
		return 166
	case metrics.MetricSourceVarnish:
		return 167
	case metrics.MetricSourceVault:
		return 168
	case metrics.MetricSourceVertica:
		return 169
	case metrics.MetricSourceVoltdb:
		return 170
	case metrics.MetricSourceVsphere:
		return 171
	case metrics.MetricSourceWeblogic:
		return 172
	case metrics.MetricSourceWin32EventLog:
		return 173
	case metrics.MetricSourceWindowsPerformanceCounters:
		return 174
	case metrics.MetricSourceWindowsService:
		return 175
	case metrics.MetricSourceWmiCheck:
		return 176
	case metrics.MetricSourceYarn:
		return 177
	case metrics.MetricSourceZk:
		return 178
	case metrics.MetricSourceContainer:
		return 180
	case metrics.MetricSourceContainerd:
		return 181
	case metrics.MetricSourceCri:
		return 182
	case metrics.MetricSourceDocker:
		return 183
	case metrics.MetricSourceNTP:
		return 184
	case metrics.MetricSourceSystemd:
		return 185
	case metrics.MetricSourceHelm:
		return 186
	case metrics.MetricSourceKubernetesAPIServer:
		return 187
	case metrics.MetricSourceKubernetesStateCore:
		return 188
	case metrics.MetricSourceOrchestrator:
		return 189
	case metrics.MetricSourceWinproc:
		return 190
	case metrics.MetricSourceFileHandle:
		return 191
	case metrics.MetricSourceWinkmem:
		return 192
	case metrics.MetricSourceIo:
		return 193
	case metrics.MetricSourceUptime:
		return 194
	case metrics.MetricSourceSbom:
		return 195
	case metrics.MetricSourceMemory:
		return 196
	case metrics.MetricSourceTCPQueueLength:
		return 197
	case metrics.MetricSourceOomKill:
		return 198
	case metrics.MetricSourceContainerLifecycle:
		return 199
	case metrics.MetricSourceJetson:
		return 200
	case metrics.MetricSourceContainerImage:
		return 201
	case metrics.MetricSourceCPU:
		return 202
	case metrics.MetricSourceLoad:
		return 203
	case metrics.MetricSourceInternal:
		return 212
	case metrics.MetricSourceSilverstripeCMS:
		return 468
	case metrics.MetricSourceSonatypeNexus:
		return 469
	case metrics.MetricSourceAnecdote:
		return 470

	case metrics.MetricSourceOpenTelemetryCollectorUnknown:
		return 0
	case metrics.MetricSourceOpenTelemetryCollectorDockerstatsReceiver:
		return 217
	case metrics.MetricSourceOpenTelemetryCollectorElasticsearchReceiver:
		return 218
	case metrics.MetricSourceOpenTelemetryCollectorExpvarReceiver:
		return 219
	case metrics.MetricSourceOpenTelemetryCollectorFilestatsReceiver:
		return 220
	case metrics.MetricSourceOpenTelemetryCollectorFlinkmetricsReceiver:
		return 221
	case metrics.MetricSourceOpenTelemetryCollectorGitproviderReceiver:
		return 222
	case metrics.MetricSourceOpenTelemetryCollectorHaproxyReceiver:
		return 223
	case metrics.MetricSourceOpenTelemetryCollectorHostmetricsReceiver:
		return 224
	case metrics.MetricSourceOpenTelemetryCollectorHttpcheckReceiver:
		return 225
	case metrics.MetricSourceOpenTelemetryCollectorIisReceiver:
		return 226
	case metrics.MetricSourceOpenTelemetryCollectorK8sclusterReceiver:
		return 227
	case metrics.MetricSourceOpenTelemetryCollectorKafkametricsReceiver:
		return 228
	case metrics.MetricSourceOpenTelemetryCollectorKubeletstatsReceiver:
		return 229
	case metrics.MetricSourceOpenTelemetryCollectorMemcachedReceiver:
		return 230
	case metrics.MetricSourceOpenTelemetryCollectorMongodbatlasReceiver:
		return 231
	case metrics.MetricSourceOpenTelemetryCollectorMongodbReceiver:
		return 232
	case metrics.MetricSourceOpenTelemetryCollectorMysqlReceiver:
		return 233
	case metrics.MetricSourceOpenTelemetryCollectorNginxReceiver:
		return 234
	case metrics.MetricSourceOpenTelemetryCollectorNsxtReceiver:
		return 235
	case metrics.MetricSourceOpenTelemetryCollectorOracledbReceiver:
		return 236
	case metrics.MetricSourceOpenTelemetryCollectorPodmanReceiver:
		return 521
	case metrics.MetricSourceOpenTelemetryCollectorPostgresqlReceiver:
		return 237
	case metrics.MetricSourceOpenTelemetryCollectorPrometheusReceiver:
		return 238
	case metrics.MetricSourceOpenTelemetryCollectorRabbitmqReceiver:
		return 239
	case metrics.MetricSourceOpenTelemetryCollectorRedisReceiver:
		return 240
	case metrics.MetricSourceOpenTelemetryCollectorRiakReceiver:
		return 241
	case metrics.MetricSourceOpenTelemetryCollectorSaphanaReceiver:
		return 242
	case metrics.MetricSourceOpenTelemetryCollectorSnmpReceiver:
		return 243
	case metrics.MetricSourceOpenTelemetryCollectorSnowflakeReceiver:
		return 244
	case metrics.MetricSourceOpenTelemetryCollectorSplunkenterpriseReceiver:
		return 245
	case metrics.MetricSourceOpenTelemetryCollectorSqlserverReceiver:
		return 246
	case metrics.MetricSourceOpenTelemetryCollectorSshcheckReceiver:
		return 247
	case metrics.MetricSourceOpenTelemetryCollectorStatsdReceiver:
		return 248
	case metrics.MetricSourceOpenTelemetryCollectorVcenterReceiver:
		return 249
	case metrics.MetricSourceOpenTelemetryCollectorZookeeperReceiver:
		return 250
	case metrics.MetricSourceOpenTelemetryCollectorActiveDirectorydsReceiver:
		return 251
	case metrics.MetricSourceOpenTelemetryCollectorAerospikeReceiver:
		return 252
	case metrics.MetricSourceOpenTelemetryCollectorApacheReceiver:
		return 253
	case metrics.MetricSourceOpenTelemetryCollectorApachesparkReceiver:
		return 254
	case metrics.MetricSourceOpenTelemetryCollectorAzuremonitorReceiver:
		return 255
	case metrics.MetricSourceOpenTelemetryCollectorBigipReceiver:
		return 256
	case metrics.MetricSourceOpenTelemetryCollectorChronyReceiver:
		return 257
	case metrics.MetricSourceOpenTelemetryCollectorCouchdbReceiver:
		return 258

	case metrics.MetricSourceArgoRollouts:
		return 314
	case metrics.MetricSourceArgoWorkflows:
		return 315
	case metrics.MetricSourceCloudera:
		return 316
	case metrics.MetricSourceDatadogClusterAgent:
		return 317
	case metrics.MetricSourceDcgm:
		return 318
	case metrics.MetricSourceEsxi:
		return 319
	case metrics.MetricSourceFluxcd:
		return 320
	case metrics.MetricSourceKarpenter:
		return 321
	case metrics.MetricSourceNvidiaTriton:
		return 322
	case metrics.MetricSourceRay:
		return 323
	case metrics.MetricSourceStrimzi:
		return 324
	case metrics.MetricSourceTekton:
		return 325
	case metrics.MetricSourceTeleport:
		return 326
	case metrics.MetricSourceTemporal:
		return 327
	case metrics.MetricSourceTorchserve:
		return 328
	case metrics.MetricSourceWeaviate:
		return 329
	case metrics.MetricSourceTraefikMesh:
		return 330
	case metrics.MetricSourceKubernetesClusterAutoscaler:
		return 331
	case metrics.MetricSourceAqua:
		return 332
	case metrics.MetricSourceAwsPricing:
		return 333
	case metrics.MetricSourceBind9:
		return 334
	case metrics.MetricSourceCfssl:
		return 335
	case metrics.MetricSourceCloudnatix:
		return 336
	case metrics.MetricSourceCloudsmith:
		return 337
	case metrics.MetricSourceCybersixgillActionableAlerts:
		return 338
	case metrics.MetricSourceCyral:
		return 339
	case metrics.MetricSourceEmqx:
		return 340
	case metrics.MetricSourceEventstore:
		return 341
	case metrics.MetricSourceExim:
		return 342
	case metrics.MetricSourceFiddler:
		return 343
	case metrics.MetricSourceFilebeat:
		return 344
	case metrics.MetricSourceFilemage:
		return 345
	case metrics.MetricSourceFluentbit:
		return 346
	case metrics.MetricSourceGatekeeper:
		return 347
	case metrics.MetricSourceGitea:
		return 348
	case metrics.MetricSourceGnatsd:
		return 349
	case metrics.MetricSourceGnatsdStreaming:
		return 350
	case metrics.MetricSourceGoPprofScraper:
		return 351
	case metrics.MetricSourceGrpcCheck:
		return 352
	case metrics.MetricSourceHikaricp:
		return 353
	case metrics.MetricSourceJfrogPlatformSelfHosted:
		return 354
	case metrics.MetricSourceKernelcare:
		return 355
	case metrics.MetricSourceLighthouse:
		return 356
	case metrics.MetricSourceLogstash:
		return 357
	case metrics.MetricSourceMergify:
		return 358
	case metrics.MetricSourceNeo4j:
		return 359
	case metrics.MetricSourceNeutrona:
		return 360
	case metrics.MetricSourceNextcloud:
		return 361
	case metrics.MetricSourceNnSdwan:
		return 362
	case metrics.MetricSourceNs1:
		return 363
	case metrics.MetricSourceNvml:
		return 364
	case metrics.MetricSourceOctoprint:
		return 365
	case metrics.MetricSourceOpenPolicyAgent:
		return 366
	case metrics.MetricSourcePhpApcu:
		return 367
	case metrics.MetricSourcePhpOpcache:
		return 368
	case metrics.MetricSourcePihole:
		return 369
	case metrics.MetricSourcePing:
		return 370
	case metrics.MetricSourcePortworx:
		return 371
	case metrics.MetricSourcePuma:
		return 372
	case metrics.MetricSourcePurefa:
		return 373
	case metrics.MetricSourcePurefb:
		return 374
	case metrics.MetricSourceRadarr:
		return 375
	case metrics.MetricSourceRebootRequired:
		return 376
	case metrics.MetricSourceRedisSentinel:
		return 377
	case metrics.MetricSourceRedisenterprise:
		return 378
	case metrics.MetricSourceRedpanda:
		return 379
	case metrics.MetricSourceRiakRepl:
		return 380
	case metrics.MetricSourceScalr:
		return 381
	case metrics.MetricSourceSendmail:
		return 382
	case metrics.MetricSourceSnmpwalk:
		return 383
	case metrics.MetricSourceSonarr:
		return 384
	case metrics.MetricSourceSortdb:
		return 385
	case metrics.MetricSourceSpeedtest:
		return 386
	case metrics.MetricSourceStardog:
		return 387
	case metrics.MetricSourceStorm:
		return 388
	case metrics.MetricSourceSyncthing:
		return 389
	case metrics.MetricSourceTidb:
		return 390
	case metrics.MetricSourceTraefik:
		return 391
	case metrics.MetricSourceUnbound:
		return 392
	case metrics.MetricSourceUnifiConsole:
		return 393
	case metrics.MetricSourceUpboundUxp:
		return 394
	case metrics.MetricSourceUpsc:
		return 395
	case metrics.MetricSourceVespa:
		return 396
	case metrics.MetricSourceWayfinder:
		return 397
	case metrics.MetricSourceZabbix:
		return 398
	case metrics.MetricSourceZenohRouter:
		return 399
	case metrics.MetricSourceVllm:
		return 412
	case metrics.MetricSourceAwsNeuron:
		return 413
	case metrics.MetricSourceAnyscale:
		return 414
	case metrics.MetricSourceAppgateSDP:
		return 415
	case metrics.MetricSourceKubeflow:
		return 416
	case metrics.MetricSourceSlurm:
		return 417
	case metrics.MetricSourceKyverno:
		return 418
	case metrics.MetricSourceTibcoEMS:
		return 419
	case metrics.MetricSourceDuckdb:
		return 423
	case metrics.MetricSourceKeda:
		return 424
	case metrics.MetricSourceMilvus:
		return 425
	case metrics.MetricSourceNvidiaNim:
		return 426
	case metrics.MetricSourceQuarkus:
		return 427
	case metrics.MetricSourceSupabase:
		return 428
	case metrics.MetricSourceDatadogOperator:
		return 456
	case metrics.MetricSourceVelero:
		return 458
	case metrics.MetricSourceCelery:
		return 464
	case metrics.MetricSourceInfiniband:
		return 465
	case metrics.MetricSourceAzureContainerAppCustom,
		metrics.MetricSourceAzureAppServiceCustom,
		metrics.MetricSourceGoogleCloudRunCustom:
		return 472
	case metrics.MetricSourceAzureContainerAppEnhanced,
		metrics.MetricSourceAzureAppServiceEnhanced,
		metrics.MetricSourceGoogleCloudRunEnhanced:
		return 473
	case metrics.MetricSourceAzureContainerAppRuntime,
		metrics.MetricSourceAzureAppServiceRuntime,
		metrics.MetricSourceGoogleCloudRunRuntime:
		return 474
	case metrics.MetricSourceWlan:
		return 475
	case metrics.MetricSourceAltairPBSPro:
		return 476
	case metrics.MetricSourceFalco:
		return 477
	case metrics.MetricSourceKrakenD:
		return 478
	case metrics.MetricSourceKuma:
		return 479
	case metrics.MetricSourceLiteLLM:
		return 480
	case metrics.MetricSourceLustre:
		return 481
	case metrics.MetricSourceProxmox:
		return 482
	case metrics.MetricSourceResilience4j:
		return 483
	case metrics.MetricSourceWindowsCertificateStore:
		return 484
	case metrics.MetricSourceBentoMl:
		return 494
	case metrics.MetricSourceHuggingFaceTgi:
		return 495
	case metrics.MetricSourceIbmSpectrumLsf:
		return 496
	case metrics.MetricSourceControlM:
		return 500
	case metrics.MetricSourceN8N:
		return 501
	case metrics.MetricSourceNutanix:
		return 502
	case metrics.MetricSourcePaloAltoPanorama:
		return 503
	case metrics.MetricSourcePrefect:
		return 504
	case metrics.MetricSourceBattery:
		return 511
	case metrics.MetricSourcePinot:
		return 512
	case metrics.MetricSourceDellPowerFlex:
		return 514
	case metrics.MetricSourceHPEArubaEdgeConnect:
		return 515
	case metrics.MetricSourceNiFi:
		return 516
	case metrics.MetricSourceCiscoSdwan:
		return 517
	case metrics.MetricSourceVersa:
		return 519
	default:
		return 0
	}

}

func originServiceToMetricSource(originService int32) metrics.MetricSource {
	// These constants map from specific fields in the 'OriginService' enum in origin.proto
	// back to MetricSource values. This is the inverse of metricSourceToOriginService.
	// Note: Multiple MetricSource values can map to the same originService value.
	// In such cases, we return a canonical/representative value.
	switch originService {
	case 0:
		return metrics.MetricSourceUnknown
	case 9:
		return metrics.MetricSourceJmxCustom
	case 10:
		return metrics.MetricSourceActiveDirectory
	case 11:
		return metrics.MetricSourceActivemqXML
	case 12:
		return metrics.MetricSourceActivemq
	case 13:
		return metrics.MetricSourceAerospike
	case 14:
		return metrics.MetricSourceAirflow
	case 15:
		return metrics.MetricSourceAmazonMsk
	case 16:
		return metrics.MetricSourceAmbari
	case 17:
		return metrics.MetricSourceApache
	case 18:
		return metrics.MetricSourceArangodb
	case 19:
		return metrics.MetricSourceArgocd
	case 20:
		return metrics.MetricSourceAspdotnet
	case 21:
		return metrics.MetricSourceAviVantage
	case 22:
		return metrics.MetricSourceAzureIotEdge
	case 23:
		return metrics.MetricSourceBoundary
	case 24:
		return metrics.MetricSourceBtrfs
	case 25:
		return metrics.MetricSourceCacti
	case 26:
		return metrics.MetricSourceCalico
	case 27:
		return metrics.MetricSourceCassandraNodetool
	case 28:
		return metrics.MetricSourceCassandra
	case 29:
		return metrics.MetricSourceCeph
	case 30:
		return metrics.MetricSourceCertManager
	case 34:
		return metrics.MetricSourceCilium
	case 36:
		return metrics.MetricSourceCitrixHypervisor
	case 37:
		return metrics.MetricSourceClickhouse
	case 38:
		return metrics.MetricSourceCloudFoundryAPI
	case 39:
		return metrics.MetricSourceCockroachdb
	case 40:
		return metrics.MetricSourceConfluentPlatform
	case 41:
		return metrics.MetricSourceConsul
	case 42:
		return metrics.MetricSourceCoredns
	case 43:
		return metrics.MetricSourceCouch
	case 44:
		return metrics.MetricSourceCouchbase
	case 45:
		return metrics.MetricSourceCrio
	case 47:
		return metrics.MetricSourceDirectory
	case 48:
		return metrics.MetricSourceDisk
	case 49:
		return metrics.MetricSourceDNSCheck
	case 50:
		return metrics.MetricSourceDotnetclr
	case 51:
		return metrics.MetricSourceDruid
	case 52:
		return metrics.MetricSourceEcsFargate
	case 53:
		return metrics.MetricSourceEksFargate
	case 54:
		return metrics.MetricSourceElastic
	case 55:
		return metrics.MetricSourceEnvoy
	case 56:
		return metrics.MetricSourceEtcd
	case 57:
		return metrics.MetricSourceExchangeServer
	case 58:
		return metrics.MetricSourceExternalDNS
	case 60:
		return metrics.MetricSourceFluentd
	case 61:
		return metrics.MetricSourceFoundationdb
	case 62:
		return metrics.MetricSourceGearmand
	case 63:
		return metrics.MetricSourceGitlabRunner
	case 64:
		return metrics.MetricSourceGitlab
	case 65:
		return metrics.MetricSourceGlusterfs
	case 66:
		return metrics.MetricSourceGoExpvar
	case 67:
		return metrics.MetricSourceGunicorn
	case 68:
		return metrics.MetricSourceHaproxy
	case 69:
		return metrics.MetricSourceHarbor
	case 70:
		return metrics.MetricSourceHazelcast
	case 71:
		return metrics.MetricSourceHdfsDatanode
	case 72:
		return metrics.MetricSourceHdfsNamenode
	case 73:
		return metrics.MetricSourceHive
	case 74:
		return metrics.MetricSourceHivemq
	case 75:
		return metrics.MetricSourceHTTPCheck
	case 76:
		return metrics.MetricSourceHudi
	case 77:
		return metrics.MetricSourceHyperv
	case 78:
		return metrics.MetricSourceIbmAce
	case 79:
		return metrics.MetricSourceIbmDb2
	case 80:
		return metrics.MetricSourceIbmI
	case 81:
		return metrics.MetricSourceIbmMq
	case 82:
		return metrics.MetricSourceIbmWas
	case 83:
		return metrics.MetricSourceIgnite
	case 84:
		return metrics.MetricSourceIis
	case 85:
		return metrics.MetricSourceImpala
	case 86:
		return metrics.MetricSourceIstio
	case 87:
		return metrics.MetricSourceJbossWildfly
	case 89:
		return metrics.MetricSourceKafkaConsumer
	case 90:
		return metrics.MetricSourceKafka
	case 91:
		return metrics.MetricSourceKong
	case 92:
		return metrics.MetricSourceKubeAPIserverMetrics
	case 93:
		return metrics.MetricSourceKubeControllerManager
	case 94:
		return metrics.MetricSourceKubeDNS
	case 95:
		return metrics.MetricSourceKubeMetricsServer
	case 96:
		return metrics.MetricSourceKubeProxy
	case 97:
		return metrics.MetricSourceKubeScheduler
	case 98:
		return metrics.MetricSourceKubelet
	case 99:
		return metrics.MetricSourceKubernetesState
	case 100:
		return metrics.MetricSourceKyototycoon
	case 101:
		return metrics.MetricSourceLighttpd
	case 102:
		return metrics.MetricSourceLinkerd
	case 103:
		return metrics.MetricSourceLinuxProcExtras
	case 104:
		return metrics.MetricSourceMapr
	case 105:
		return metrics.MetricSourceMapreduce
	case 106:
		return metrics.MetricSourceMarathon
	case 107:
		return metrics.MetricSourceMarklogic
	case 108:
		return metrics.MetricSourceMcache
	case 109:
		return metrics.MetricSourceMesosMaster
	case 110:
		return metrics.MetricSourceMesosSlave
	case 111:
		return metrics.MetricSourceMongo
	case 112:
		return metrics.MetricSourceMysql
	case 113:
		return metrics.MetricSourceNagios
	case 114:
		return metrics.MetricSourceNetwork
	case 115:
		return metrics.MetricSourceNfsstat
	case 116:
		return metrics.MetricSourceNginxIngressController
	case 117:
		return metrics.MetricSourceNginx
	case 118:
		return metrics.MetricSourceOpenldap
	case 119:
		return metrics.MetricSourceOpenmetrics
	case 120:
		return metrics.MetricSourceOpenstackController
	case 121:
		return metrics.MetricSourceOpenstack
	case 122:
		return metrics.MetricSourceOracle
	case 124:
		return metrics.MetricSourcePdhCheck
	case 125:
		return metrics.MetricSourcePgbouncer
	case 126:
		return metrics.MetricSourcePhpFpm
	case 127:
		return metrics.MetricSourcePostfix
	case 128:
		return metrics.MetricSourcePostgres
	case 129:
		return metrics.MetricSourcePowerdnsRecursor
	case 130:
		return metrics.MetricSourcePresto
	case 131:
		return metrics.MetricSourceProcess
	case 132:
		return metrics.MetricSourcePrometheus
	case 133:
		return metrics.MetricSourceProxysql
	case 134:
		return metrics.MetricSourcePulsar
	case 135:
		return metrics.MetricSourceRabbitmq
	case 136:
		return metrics.MetricSourceRedisdb
	case 137:
		return metrics.MetricSourceRethinkdb
	case 138:
		return metrics.MetricSourceRiak
	case 139:
		return metrics.MetricSourceRiakcs
	case 140:
		return metrics.MetricSourceSapHana
	case 141:
		return metrics.MetricSourceScylla
	case 143:
		return metrics.MetricSourceSilk
	case 144:
		return metrics.MetricSourceSinglestore
	case 145:
		return metrics.MetricSourceSnmp
	case 146:
		return metrics.MetricSourceSnowflake
	case 147:
		return metrics.MetricSourceSolr
	case 148:
		return metrics.MetricSourceSonarqube
	case 149:
		return metrics.MetricSourceSpark
	case 150:
		return metrics.MetricSourceSqlserver
	case 151:
		return metrics.MetricSourceSquid
	case 152:
		return metrics.MetricSourceSSHCheck
	case 153:
		return metrics.MetricSourceStatsd
	case 154:
		return metrics.MetricSourceSupervisord
	case 155:
		return metrics.MetricSourceSystemCore
	case 156:
		return metrics.MetricSourceSystemSwap
	case 157:
		return metrics.MetricSourceTCPCheck
	case 158:
		return metrics.MetricSourceTeamcity
	case 160:
		return metrics.MetricSourceTeradata
	case 161:
		return metrics.MetricSourceTLS
	case 162:
		return metrics.MetricSourceTokumx
	case 163:
		return metrics.MetricSourceTomcat
	case 164:
		return metrics.MetricSourceTrafficServer
	case 165:
		return metrics.MetricSourceTwemproxy
	case 166:
		return metrics.MetricSourceTwistlock
	case 167:
		return metrics.MetricSourceVarnish
	case 168:
		return metrics.MetricSourceVault
	case 169:
		return metrics.MetricSourceVertica
	case 170:
		return metrics.MetricSourceVoltdb
	case 171:
		return metrics.MetricSourceVsphere
	case 172:
		return metrics.MetricSourceWeblogic
	case 173:
		return metrics.MetricSourceWin32EventLog
	case 174:
		return metrics.MetricSourceWindowsPerformanceCounters
	case 175:
		return metrics.MetricSourceWindowsService
	case 176:
		return metrics.MetricSourceWmiCheck
	case 177:
		return metrics.MetricSourceYarn
	case 178:
		return metrics.MetricSourceZk
	case 180:
		return metrics.MetricSourceContainer
	case 181:
		return metrics.MetricSourceContainerd
	case 182:
		return metrics.MetricSourceCri
	case 183:
		return metrics.MetricSourceDocker
	case 184:
		return metrics.MetricSourceNTP
	case 185:
		return metrics.MetricSourceSystemd
	case 186:
		return metrics.MetricSourceHelm
	case 187:
		return metrics.MetricSourceKubernetesAPIServer
	case 188:
		return metrics.MetricSourceKubernetesStateCore
	case 189:
		return metrics.MetricSourceOrchestrator
	case 190:
		return metrics.MetricSourceWinproc
	case 191:
		return metrics.MetricSourceFileHandle
	case 192:
		return metrics.MetricSourceWinkmem
	case 193:
		return metrics.MetricSourceIo
	case 194:
		return metrics.MetricSourceUptime
	case 195:
		return metrics.MetricSourceSbom
	case 196:
		return metrics.MetricSourceMemory
	case 197:
		return metrics.MetricSourceTCPQueueLength
	case 198:
		return metrics.MetricSourceOomKill
	case 199:
		return metrics.MetricSourceContainerLifecycle
	case 200:
		return metrics.MetricSourceJetson
	case 201:
		return metrics.MetricSourceContainerImage
	case 202:
		return metrics.MetricSourceCPU
	case 203:
		return metrics.MetricSourceLoad
	case 212:
		return metrics.MetricSourceInternal
	case 217:
		return metrics.MetricSourceOpenTelemetryCollectorDockerstatsReceiver
	case 218:
		return metrics.MetricSourceOpenTelemetryCollectorElasticsearchReceiver
	case 219:
		return metrics.MetricSourceOpenTelemetryCollectorExpvarReceiver
	case 220:
		return metrics.MetricSourceOpenTelemetryCollectorFilestatsReceiver
	case 221:
		return metrics.MetricSourceOpenTelemetryCollectorFlinkmetricsReceiver
	case 222:
		return metrics.MetricSourceOpenTelemetryCollectorGitproviderReceiver
	case 223:
		return metrics.MetricSourceOpenTelemetryCollectorHaproxyReceiver
	case 224:
		return metrics.MetricSourceOpenTelemetryCollectorHostmetricsReceiver
	case 225:
		return metrics.MetricSourceOpenTelemetryCollectorHttpcheckReceiver
	case 226:
		return metrics.MetricSourceOpenTelemetryCollectorIisReceiver
	case 227:
		return metrics.MetricSourceOpenTelemetryCollectorK8sclusterReceiver
	case 228:
		return metrics.MetricSourceOpenTelemetryCollectorKafkametricsReceiver
	case 229:
		return metrics.MetricSourceOpenTelemetryCollectorKubeletstatsReceiver
	case 230:
		return metrics.MetricSourceOpenTelemetryCollectorMemcachedReceiver
	case 231:
		return metrics.MetricSourceOpenTelemetryCollectorMongodbatlasReceiver
	case 232:
		return metrics.MetricSourceOpenTelemetryCollectorMongodbReceiver
	case 233:
		return metrics.MetricSourceOpenTelemetryCollectorMysqlReceiver
	case 234:
		return metrics.MetricSourceOpenTelemetryCollectorNginxReceiver
	case 235:
		return metrics.MetricSourceOpenTelemetryCollectorNsxtReceiver
	case 236:
		return metrics.MetricSourceOpenTelemetryCollectorOracledbReceiver
	case 237:
		return metrics.MetricSourceOpenTelemetryCollectorPostgresqlReceiver
	case 238:
		return metrics.MetricSourceOpenTelemetryCollectorPrometheusReceiver
	case 239:
		return metrics.MetricSourceOpenTelemetryCollectorRabbitmqReceiver
	case 240:
		return metrics.MetricSourceOpenTelemetryCollectorRedisReceiver
	case 241:
		return metrics.MetricSourceOpenTelemetryCollectorRiakReceiver
	case 242:
		return metrics.MetricSourceOpenTelemetryCollectorSaphanaReceiver
	case 243:
		return metrics.MetricSourceOpenTelemetryCollectorSnmpReceiver
	case 244:
		return metrics.MetricSourceOpenTelemetryCollectorSnowflakeReceiver
	case 245:
		return metrics.MetricSourceOpenTelemetryCollectorSplunkenterpriseReceiver
	case 246:
		return metrics.MetricSourceOpenTelemetryCollectorSqlserverReceiver
	case 247:
		return metrics.MetricSourceOpenTelemetryCollectorSshcheckReceiver
	case 248:
		return metrics.MetricSourceOpenTelemetryCollectorStatsdReceiver
	case 249:
		return metrics.MetricSourceOpenTelemetryCollectorVcenterReceiver
	case 250:
		return metrics.MetricSourceOpenTelemetryCollectorZookeeperReceiver
	case 251:
		return metrics.MetricSourceOpenTelemetryCollectorActiveDirectorydsReceiver
	case 252:
		return metrics.MetricSourceOpenTelemetryCollectorAerospikeReceiver
	case 253:
		return metrics.MetricSourceOpenTelemetryCollectorApacheReceiver
	case 254:
		return metrics.MetricSourceOpenTelemetryCollectorApachesparkReceiver
	case 255:
		return metrics.MetricSourceOpenTelemetryCollectorAzuremonitorReceiver
	case 256:
		return metrics.MetricSourceOpenTelemetryCollectorBigipReceiver
	case 257:
		return metrics.MetricSourceOpenTelemetryCollectorChronyReceiver
	case 258:
		return metrics.MetricSourceOpenTelemetryCollectorCouchdbReceiver
	case 314:
		return metrics.MetricSourceArgoRollouts
	case 315:
		return metrics.MetricSourceArgoWorkflows
	case 316:
		return metrics.MetricSourceCloudera
	case 317:
		return metrics.MetricSourceDatadogClusterAgent
	case 318:
		return metrics.MetricSourceDcgm
	case 319:
		return metrics.MetricSourceEsxi
	case 320:
		return metrics.MetricSourceFluxcd
	case 321:
		return metrics.MetricSourceKarpenter
	case 322:
		return metrics.MetricSourceNvidiaTriton
	case 323:
		return metrics.MetricSourceRay
	case 324:
		return metrics.MetricSourceStrimzi
	case 325:
		return metrics.MetricSourceTekton
	case 326:
		return metrics.MetricSourceTeleport
	case 327:
		return metrics.MetricSourceTemporal
	case 328:
		return metrics.MetricSourceTorchserve
	case 329:
		return metrics.MetricSourceWeaviate
	case 330:
		return metrics.MetricSourceTraefikMesh
	case 331:
		return metrics.MetricSourceKubernetesClusterAutoscaler
	case 332:
		return metrics.MetricSourceAqua
	case 333:
		return metrics.MetricSourceAwsPricing
	case 334:
		return metrics.MetricSourceBind9
	case 335:
		return metrics.MetricSourceCfssl
	case 336:
		return metrics.MetricSourceCloudnatix
	case 337:
		return metrics.MetricSourceCloudsmith
	case 338:
		return metrics.MetricSourceCybersixgillActionableAlerts
	case 339:
		return metrics.MetricSourceCyral
	case 340:
		return metrics.MetricSourceEmqx
	case 341:
		return metrics.MetricSourceEventstore
	case 342:
		return metrics.MetricSourceExim
	case 343:
		return metrics.MetricSourceFiddler
	case 344:
		return metrics.MetricSourceFilebeat
	case 345:
		return metrics.MetricSourceFilemage
	case 346:
		return metrics.MetricSourceFluentbit
	case 347:
		return metrics.MetricSourceGatekeeper
	case 348:
		return metrics.MetricSourceGitea
	case 349:
		return metrics.MetricSourceGnatsd
	case 350:
		return metrics.MetricSourceGnatsdStreaming
	case 351:
		return metrics.MetricSourceGoPprofScraper
	case 352:
		return metrics.MetricSourceGrpcCheck
	case 353:
		return metrics.MetricSourceHikaricp
	case 354:
		return metrics.MetricSourceJfrogPlatformSelfHosted
	case 355:
		return metrics.MetricSourceKernelcare
	case 356:
		return metrics.MetricSourceLighthouse
	case 357:
		return metrics.MetricSourceLogstash
	case 358:
		return metrics.MetricSourceMergify
	case 359:
		return metrics.MetricSourceNeo4j
	case 360:
		return metrics.MetricSourceNeutrona
	case 361:
		return metrics.MetricSourceNextcloud
	case 362:
		return metrics.MetricSourceNnSdwan
	case 363:
		return metrics.MetricSourceNs1
	case 364:
		return metrics.MetricSourceNvml
	case 365:
		return metrics.MetricSourceOctoprint
	case 366:
		return metrics.MetricSourceOpenPolicyAgent
	case 367:
		return metrics.MetricSourcePhpApcu
	case 368:
		return metrics.MetricSourcePhpOpcache
	case 369:
		return metrics.MetricSourcePihole
	case 370:
		return metrics.MetricSourcePing
	case 371:
		return metrics.MetricSourcePortworx
	case 372:
		return metrics.MetricSourcePuma
	case 373:
		return metrics.MetricSourcePurefa
	case 374:
		return metrics.MetricSourcePurefb
	case 375:
		return metrics.MetricSourceRadarr
	case 376:
		return metrics.MetricSourceRebootRequired
	case 377:
		return metrics.MetricSourceRedisSentinel
	case 378:
		return metrics.MetricSourceRedisenterprise
	case 379:
		return metrics.MetricSourceRedpanda
	case 380:
		return metrics.MetricSourceRiakRepl
	case 381:
		return metrics.MetricSourceScalr
	case 382:
		return metrics.MetricSourceSendmail
	case 383:
		return metrics.MetricSourceSnmpwalk
	case 384:
		return metrics.MetricSourceSonarr
	case 385:
		return metrics.MetricSourceSortdb
	case 386:
		return metrics.MetricSourceSpeedtest
	case 387:
		return metrics.MetricSourceStardog
	case 388:
		return metrics.MetricSourceStorm
	case 389:
		return metrics.MetricSourceSyncthing
	case 390:
		return metrics.MetricSourceTidb
	case 391:
		return metrics.MetricSourceTraefik
	case 392:
		return metrics.MetricSourceUnbound
	case 393:
		return metrics.MetricSourceUnifiConsole
	case 394:
		return metrics.MetricSourceUpboundUxp
	case 395:
		return metrics.MetricSourceUpsc
	case 396:
		return metrics.MetricSourceVespa
	case 397:
		return metrics.MetricSourceWayfinder
	case 398:
		return metrics.MetricSourceZabbix
	case 399:
		return metrics.MetricSourceZenohRouter
	case 412:
		return metrics.MetricSourceVllm
	case 413:
		return metrics.MetricSourceAwsNeuron
	case 414:
		return metrics.MetricSourceAnyscale
	case 415:
		return metrics.MetricSourceAppgateSDP
	case 416:
		return metrics.MetricSourceKubeflow
	case 417:
		return metrics.MetricSourceSlurm
	case 418:
		return metrics.MetricSourceKyverno
	case 419:
		return metrics.MetricSourceTibcoEMS
	case 423:
		return metrics.MetricSourceDuckdb
	case 424:
		return metrics.MetricSourceKeda
	case 425:
		return metrics.MetricSourceMilvus
	case 426:
		return metrics.MetricSourceNvidiaNim
	case 427:
		return metrics.MetricSourceQuarkus
	case 428:
		return metrics.MetricSourceSupabase
	case 430:
		return metrics.MetricSourceFlyIo
	case 431:
		return metrics.MetricSourceKepler
	case 432:
		return metrics.MetricSourceOctopusDeploy
	case 433:
		return metrics.MetricSourceScaphandre
	case 436:
		return metrics.MetricSourceJenkins
	case 437:
		return metrics.MetricSourceKubeVirtAPI
	case 438:
		return metrics.MetricSourceKubeVirtController
	case 439:
		return metrics.MetricSourceKubeVirtHandler
	case 440:
		return metrics.MetricSourceCloudFoundry
	case 456:
		return metrics.MetricSourceDatadogOperator
	case 458:
		return metrics.MetricSourceVelero
	case 464:
		return metrics.MetricSourceCelery
	case 465:
		return metrics.MetricSourceInfiniband
	case 466:
		return metrics.MetricSourceGPU
	case 468:
		return metrics.MetricSourceSilverstripeCMS
	case 469:
		return metrics.MetricSourceSonatypeNexus
	case 470:
		return metrics.MetricSourceAnecdote
	case 472:
		// Multiple sources map to 472: AzureContainerAppCustom, AzureAppServiceCustom, GoogleCloudRunCustom
		return metrics.MetricSourceAzureContainerAppCustom
	case 473:
		// Multiple sources map to 473: AzureContainerAppEnhanced, AzureAppServiceEnhanced, GoogleCloudRunEnhanced
		return metrics.MetricSourceAzureContainerAppEnhanced
	case 474:
		// Multiple sources map to 474: AzureContainerAppRuntime, AzureAppServiceRuntime, GoogleCloudRunRuntime
		return metrics.MetricSourceAzureContainerAppRuntime
	case 475:
		return metrics.MetricSourceWlan
	case 476:
		return metrics.MetricSourceAltairPBSPro
	case 477:
		return metrics.MetricSourceFalco
	case 478:
		return metrics.MetricSourceKrakenD
	case 479:
		return metrics.MetricSourceKuma
	case 480:
		return metrics.MetricSourceLiteLLM
	case 481:
		return metrics.MetricSourceLustre
	case 482:
		return metrics.MetricSourceProxmox
	case 483:
		return metrics.MetricSourceResilience4j
	case 484:
		return metrics.MetricSourceWindowsCertificateStore
	case 494:
		return metrics.MetricSourceBentoMl
	case 495:
		return metrics.MetricSourceHuggingFaceTgi
	case 496:
		return metrics.MetricSourceIbmSpectrumLsf
	case 500:
		return metrics.MetricSourceControlM
	case 501:
		return metrics.MetricSourceN8N
	case 502:
		return metrics.MetricSourceNutanix
	case 503:
		return metrics.MetricSourcePaloAltoPanorama
	case 504:
		return metrics.MetricSourcePrefect
	case 511:
		return metrics.MetricSourceBattery
	case 512:
		return metrics.MetricSourcePinot
	case 514:
		return metrics.MetricSourceDellPowerFlex
	case 515:
		return metrics.MetricSourceHPEArubaEdgeConnect
	case 516:
		return metrics.MetricSourceNiFi
	case 517:
		return metrics.MetricSourceCiscoSdwan
	case 519:
		return metrics.MetricSourceVersa
	case 521:
		return metrics.MetricSourceOpenTelemetryCollectorPodmanReceiver
	default:
		return metrics.MetricSourceUnknown
	}
}
