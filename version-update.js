const fs = require('fs');

const preCommit = (props) => {
    const chartPath = 'helm/publicapi/Chart.yaml';
    fs.writeFileSync('version.txt', props.version);
    console.log('Wrote version.txt');
    const chart = fs.readFileSync(chartPath);
    const chartNew = chart.toString()
        .replace(/^version: .*$/mg, `version: ${props.version}`)
        .replace(/^appVersion: .*$/mg, `appVersion: ${props.version}`);
    fs.writeFileSync(chartPath, chartNew);
    console.log('Wrote ', chartPath);
}

exports.preCommit = preCommit;

// preCommit({version: 'test-ver'});
