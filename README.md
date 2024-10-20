<p align="center">
  <img alt="3D pixelated Kibisis icon" src=".github/assets/logo@405x128.png" style="padding-top: 15px" height="128" />
</p>

<h1 align="center">
   Kibisis API
</h1>

<p align="center">
  The Kibisis API is a set of DigitalOcean functions that act as a RESTful API for the Kibisis ecosystem.
</p>

<p align="center">
  <a href="https://github.com/kibis-is/api/releases/latest">
    <img alt="GitHub release" src="https://img.shields.io/github/v/release/kibis-is/api?&logo=github">
  </a>
  <a href="https://github.com/kibis-is/api/releases/latest">
    <img alt="GitHub release date - published at" src="https://img.shields.io/github/release-date/kibis-is/api?logo=github">
  </a>
</p>

<p align="center">
  <a href="https://github.com/kibis-is/api/releases">
    <img alt="GitHub Pre-release" src="https://img.shields.io/github/v/release/kibis-is/api?include_prereleases&label=pre-release&logo=github">
  </a>
  <a href="https://github.com/kibis-is/api/releases">
    <img alt="GitHub Pre-release Date - Published At" src="https://img.shields.io/github/release-date-pre/kibis-is/api?label=pre-release date&logo=github">
  </a>
</p>

<p align="center">
  <a href="https://github.com/kibis-is/api/blob/main/COPYING">
    <img alt="GitHub license" src="https://img.shields.io/github/license/kibis-is/api">
  </a>
</p>

#### Table Of Contents

* [1. Overview](#-1-overview)
  - [1.1. Project Structure](#11-project-structure)
* [2. Development](#-2-development)
  - [2.1. Requirements](#21-requirements)
  - [2.2. Create Personal Doppler Config](#22-create-personal-doppler-config)
  - [2.3. Setup `doppler`](#23-setup-doppler)
  - [2.4. Setup `doctl`](#24-setup-doctl)
  - [2.5. Create A Personal Namespace (optional)](#25-create-a-personal-namespace-optional)
  - [2.6. Deploy The Functions To The Namespace](#26-deploy-the-functions-to-the-namespace)
* [3. Appendix](#-3-appendix)
  - [3.1. Useful Commands](#31-useful-commands)
* [4. How To Contribute](#-4-how-to-contribute)
* [5. License](#-5-license)

## 🗂️ 1. Overview

### 1.1. Project Structure

The project structure is based on the DigitalOcean [functions](https://docs.digitalocean.com/products/functions/how-to/structure-projects) project structure.
However, the core directories `lib` and `packages` have specific functionality:

* `packages` - This is where each function resides. Each package/function relates to a versioned API path. For example, `v1/quests` will correspond to an API path `https://<endpoint>/v1/quests`
* `lib` - This contains independent modules that are referenced from the packages.

> ⚠️ **NOTE:** Each path must parse the request header and handle the method, i.e. GET, POST, DELETE e.t.c.

<sup>[Back to top ^][table-of-contents]</sup>

## 🛠️ 2. Development

### 2.1. Requirements

* [DigitalOcean CLI][doctl]
* [Doppler CLI][doppler]
* [Golang v1.20+][golang]
* [Make][make]

<sup>[Back to top ^][table-of-contents]</sup>

## 2.2. Create Personal Doppler Config

To start using your own Doppler config, go to the project on [Doppler](https://dashboard.doppler.com/workplace/ae8c01548486ba93b8fd/projects/api) and press the "+" to create a new personal branch config in the "Development" config

<p align="center">
  <img alt="Screen grab of the Doppler dashboard when creating a branch config" src=".github/assets/create_doppler_config.png" style="padding-top: 15px" height="512" />
</p>

> ⚠️ **NOTE:** Use your name in lowercase with underscores instead of spaces (snake_case).

<sup>[Back to top ^][table-of-contents]</sup>

### 2.3. Setup `doppler`

Once the branch project config has been setup, follow the instructions [here](https://docs.doppler.com/docs/install-cli#project-setup) to:
* login to Doppler, and;
* setup Doppler to use the `api` project with your personal config.

> ⚠️ **NOTE:** When naming your token, it is recommended you use: "<name>-<device_name>".

<sup>[Back to top ^][table-of-contents]</sup>

### 2.4. Setup `doctl`

The DigitalOcean CLI client, `doctl`, is used to deploy the function to a remote environment tha can be used to develop.

Follow the instructions outlined in the [documentation][doctl].

> ⚠️ **NOTE:** Once you have setup `doctl` make sure you also install the serverless subcommand using `doctl serverless install`

<sup>[Back to top ^][table-of-contents]</sup>

### 2.5. Create A Personal Namespace (optional)

> ⚠️ **NOTE:** If you have a personal namespace already setup, you can skip this step.

1. Setup a personal namespace using your name suffixed by the word "`namespace`" ensuring to use kebab-case:

```shell script
doctl serverless namespaces create --label="kieran-namespace" --region="ams3"
```

> ⚠️ **NOTE:** The above example shows setting up a namespace in the Amsterdam region, but it is better if you use a region that is closest to you. Run the command: `doctl serverless namespaces list-regions` to get a list of regions and replace the `--region` value with the desired region form the list.

<sup>[Back to top ^][table-of-contents]</sup>

### 2.6. Deploy The Functions To The Namespace

1. Deploy the function to the namespace that was created in the [previous](#25-create-a-personal-namespace-optional) step:
```shell
make deploy
```

2. Get the URL of the deployed function:
```shell
doctl serverless functions get <package>/<function> --url
```

This will return a URL in the form of:

```text
https://faas-ams3-2a2df116.doserverless.co/api/v1/web/<namespace>/<package>/<function>
```

Use this URL to interact with the API.

<sup>[Back to top ^][table-of-contents]</sup>

## 📑 3. Appendix

### 3.1. Useful Commands

| Command        | Description                                                                                                                             |
|----------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| `make`         | Deploys the functions to the remote namespace. Intended for development purposes only.                                                  |
| `make clean`   | Removes build files and configurations.                                                                                                 |
| `make deploy`  | Fetches secrets from Doppler and deploys the functions to the configured namespace. Intended for development purposes only.             |
| `make list`    | Lists the deployed functions in the configured namespace.                                                                               |
| `make logs`    | Outputs the activation logs for the deployed functions.                                                                                 |
| `make test`    | Runs the tests for each function.                                                                                                       |
| `make setup`   | Installs dependencies.                                                                                                                  |
| `make swagger` | Parses the functions and creates the swagger.json to `packages/vX/swagger` directories.                                                 |
| `make watch`   | Fetches secrets from Doppler and watches for code changes and redeploys functions to namespace. Intended for development purposes only. |

<sup>[Back to top ^][table-of-contents]</sup>

## 👏 4. How To Contribute

Please read the [**Contributing Guide**][contribute] to learn about the development process.

<sup>[Back to top ^][table-of-contents]</sup>

## 📄 5. License

Please refer to the [COPYING][copying] file.

<sup>[Back to top ^][table-of-contents]</sup>

<!-- Links -->
[contribute]: ./CONTRIBUTING.md
[copying]: ./COPYING
[doctl]: https://docs.digitalocean.com/reference/doctl/how-to/install/
[doppler]: https://docs.doppler.com/docs/install-cli
[golang]: https://go.dev/doc/install
[make]: https://www.gnu.org/software/make/
[table-of-contents]: #table-of-contents
