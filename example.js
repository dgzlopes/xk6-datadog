import datadog from 'k6/x/datadog';

const client = datadog.newClient("<your-api-key>", "<your-app-key>");

export default function() {
    var ingestAvgPerZone = client.metricQuery("service_ingest.kb.avg{*}by{availability-zone}", 3600);

    ingestAvgPerZone.forEach(t => {
        console.log(`${t.expression} with ${t.points.length} data points`);
        console.log(`unit=${t.unit} avg=${t.avg()} min=${t.min()} max=${t.max()}`);
    });

    var dockerThreadCount = client.metricQuery("docker.thread.count{*}");

    dockerThreadCount.forEach(t => {
        console.log(`${t.expression} with ${t.points.length} data points`);
        console.log(`unit=${t.unit} avg=${t.avg()} min=${t.min()} max=${t.max()}`);
    });
}