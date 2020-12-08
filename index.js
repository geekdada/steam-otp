#!/usr/bin/env node

/**
 * steam-otp-cli
 * Generate Steam one-time password offline.
 *
 * @author Roy Li <https://github.com/geekdada/steam-otp-cli>
 */

const inquirer = require('inquirer')
const SteamTotp = require('steam-totp')

const init = require('./utils/init')
const cli = require('./utils/cli')
const log = require('./utils/log')

const input = cli.input
const flags = cli.flags
const { debug } = flags

;(async () => {
  init()
  input.includes(`help`) && cli.showHelp(0)

  debug && log(flags)

  const answers = await inquirer.prompt([
    {
      type: 'password',
      name: 'sharedSecret',
      message: 'Input the shared secret you get from SteamDesktopAuthenticator',
      validate(input) {
        return input.length > 0
      },
    },
  ])

  try {
    const code = SteamTotp.generateAuthCode(answers.sharedSecret)

    console.log()
    console.log(code)
  } catch (err) {
    log(err)
  }
})()
