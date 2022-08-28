const fs = require('fs');

exports.preCommit = async (props) => {
    const chartPath = 'helm/publicapi/Chart.yml';
    fs.writeFileSync('version.txt', props.version);
    console.log('Wrote version.txt');
/*
    const chart = fs.readFileSync(chartPath);
    fs.writeFileSync(chartPath, chart.toString()
        .replace(/^version: .*$/, `version: ${props.version}`)
        .replace(/^appVersion: .*$/, `appVersion: ${props.version}`));
    console.log('Wrote ', chartPath);
*/
}