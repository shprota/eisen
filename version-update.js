const fs = require('fs');

module.exports = function preCommit(props) {
    const chartPath = 'helm/publicapi/Chart.yml';
    const chart = fs.readFileSync(chartPath);
    fs.writeFileSync(chartPath, chart.toString()
        .replace(/^version: .*$/, `version: ${props.version}`)
        .replace(/^appVersion: .*$/, `appVersion: ${props.version}`));
}