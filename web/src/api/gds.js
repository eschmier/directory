const api = require('./trisa/gds/api/v1beta1/api_grpc_web_pb');
const models = require('./trisa/gds/models/v1beta1/models_pb');

// TODO: configure to Proxy endpoint of GDS
const endpoint = "http://localhost:8080"


class GDS {
  constructor(endpoint) {
    this.client = new api.TRISADirectoryClient(endpoint);
  }

  verificationStatus = (num) => {
    for (const key in models.VerificationState) {
      if (num === models.VerificationState[key]) {
        return key;
      }
    };
  }

  lookup = (query, inputType) => {
    const req = new api.LookupRequest();
    switch (inputType) {
      case "uuid":
        req.setId(query);
        req.setRegisteredDirectory("vaspdirectory.net");
        break
      case "common name":
        req.setCommonName(query);
        break
      default:
        throw new Error("unacceptable input type to lookup query");
    }

    let self = this;
    return new Promise(resolve => {
      self.client.lookup(req, {}, (err, rep) => {
        if (err || !rep) {
          console.log(err);
          resolve(undefined);
        }
        resolve(rep.toObject());
      });
    });
  }

  getStatus = (commonName) => {
    const req = new api.StatusRequest();
    req.setCommonName(commonName);
    this.client.status(req, {}, (err, rep) => {
      if (err || !rep) {
        console.log(err);
        return
      }
      console.log(this.verificationStatus(rep.getVerificationStatus()));
      console.log(rep.toObject());
    });
  };

}


export default new GDS(endpoint);