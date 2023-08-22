import { StatsService } from 'app/gen/ts/api/services/stats/stats_connect';

const deps = {
    proto: {
        StatsService: StatsService,
    }
}

const requireFile = require.context(
    '../services',
    false,
    /[\w-]+\.js$/
)

const services = {}
requireFile.keys().forEach(fileName => {
    const config = requireFile(fileName)
    const name = fileName
        .replace(/^\.\//, '')
        .replace(/^\.\/_/, '')
        .replace(/\.\w+$/, '')
    const Service = config.default || config
    services[name] = new Service(deps)
})

export default ({ Vue }) => {
    Vue.prototype.$service = services
}
