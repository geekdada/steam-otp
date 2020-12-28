# steam-otp-cli

## Install

```bash
$ curl -sf https://gobinaries.com/geekdada/steam-otp | sh
```

## Usage

Get the `shared_secret` first, using [SteamDesktopAuthenticator](https://github.com/Jessecar96/SteamDesktopAuthenticator).

Once you have the code, store it in your vault (1Password or similar thing).

Run the command and input the `shared_secret`.

```bash
$ steam-otp
```

Or, using `--secret` to pass the secret (NOT SAFE).

```bash
$ steam-otp --secret [SHARED SECRET]
```

**Notice**

- Your shared secret will never be stored by `zsh_history` or `bash_history`.
- You should never save the shared secret in plain text.
- This method may violate the TOS of Steam, use at your discretion.

## License

[MIT](https://github.com/geekdada/steam-otp-cli/blob/master/LICENSE)
