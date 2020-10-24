# sifnode

**sifnode** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
# clone
git clone git@github.com:Sifchain/sifnode.git && cd sifnode
#
git branch
# build
make install
# cd
cd ./build
# Scaffold
rake 'genesis:sifnode:scaffold[monkey-bars, 190cb35265860f182e35a3bceeb297082858eebd@35.166.247.98:26656, http://35.166.247.98:26657/genesis]'
# Run
sifnoded start
```
For additional help, check out our demo video https://www.youtube.com/watch?v=1kjdjCEcYak&feature=youtu.be&ab_channel=utx0_

You can also ask questions at our Discord channel - https://discord.gg/SE2dg8

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
