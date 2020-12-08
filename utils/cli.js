const meow = require('meow')
const meowHelp = require('cli-meow-help')

const flags = {
  debug: {
    type: `boolean`,
    default: false,
    alias: `d`,
    desc: `Print debug info`,
  },
  version: {
    type: `boolean`,
    alias: `v`,
    desc: `Print CLI version`,
  },
}

const commands = {
  help: { desc: `Print help info` },
}

const helpText = meowHelp({
  name: `steam-otp`,
  flags,
  commands,
})

const options = {
  inferType: true,
  description: false,
  hardRejection: false,
  flags,
}

module.exports = meow(helpText, options)
