import { SimpleSpanProcessor } from "@opentelemetry/sdk-trace-base"
import { Resource } from "@opentelemetry/resources"
import { SemanticResourceAttributes } from "@opentelemetry/semantic-conventions"
import { NodeTracerProvider } from "@opentelemetry/sdk-trace-node"
import { diag, DiagConsoleLogger, DiagLogLevel } from "@opentelemetry/api"
import { registerInstrumentations } from "@opentelemetry/instrumentation"
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-grpc'
import { CompressionAlgorithm } from '@opentelemetry/otlp-exporter-base'
import grpc from '@grpc/grpc-js'
import { HttpInstrumentation } from "@opentelemetry/instrumentation-http"

diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.DEBUG)

registerInstrumentations({ instrumentations: [new HttpInstrumentation()] })

const resource =
  Resource.default().merge(
    new Resource({
      [SemanticResourceAttributes.SERVICE_NAME]: "stickerfy-ui",
    })
  )

const provider = new NodeTracerProvider({
  resource: resource,
})

const collectorOptions = {
  url: process.env.OTEL_COLLECTOR_OTPL_TRACES_ENDPOINT || 'http://localhost:4317',
  timeoutMillis: 10000,
  compression: CompressionAlgorithm.GZIP,
  credentials: grpc.credentials.createInsecure(),
}
const exporterGrpc = new OTLPTraceExporter(collectorOptions)

const processor = new SimpleSpanProcessor(exporterGrpc)
provider.addSpanProcessor(processor)

provider.register()
