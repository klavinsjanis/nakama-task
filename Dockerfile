FROM heroiclabs/nakama-pluginbuilder:3.21.1 AS go-builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build --trimpath --buildmode=plugin -o ./app.so

FROM registry.heroiclabs.com/heroiclabs/nakama:3.21.1

COPY --from=go-builder /app/app.so /nakama/data/modules/
COPY --from=go-builder /app/files/ /nakama/data/files/