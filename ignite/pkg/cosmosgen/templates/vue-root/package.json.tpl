{
  "name": "{{ .PackageNS }}-vuex",
  "version": "0.0.1",
  "description": "Autogenerated Vuex Stores",
  "author": "Ignite Codegen <hello@ignite.com>",
  "license": "Apache-2.0",
  "licenses": [
    {
      "type": "Apache-2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0"
    }
  ],
  "scripts": {
    "build": "tsc",
    "postinstall": "node postinstall.js",
    "prepublishOnly": "node local-check.js && tsc"
  },
  "main": "./lib/index.js",
  "publishConfig": {
    "access": "public"
  },
  "dependencies": {
    "buffer": "^6.0.3"
  },
  "devDependencies": {
    "typescript": "^4.8.4"
  },
  "peerDependencies": {
    "@cosmjs/proto-signing": "0.27.0",
    "@cosmjs/stargate": "0.27.0",
    "vue": "3.2.31"
  }
}
