<template>
  <div class="login">
    <snackbar :obj="obj" />

    <v-card
      class="grey pa-6 lighten-4 mx-auto rounded-xl"
      width="400"
      height="450"
    >
      <v-layout>
        <v-flex>
          <h1
            class="text-center mb-8 font-weight-black grey--text text-capitalize text--darken-2"
          >
            Login Page
          </h1>
        </v-flex>
      </v-layout>
      <v-layout>
        <v-flex>
          <v-text-field
            v-model="User_Details.userId"
            label="UserId"
            outlined
            type="text"
            class="mx-12 mb-5"
          ></v-text-field>
        </v-flex>
      </v-layout>
      <v-layout>
        <v-flex
          ><v-text-field
            outlined
            v-model="User_Details.password"
            :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
            :rules="[rules.required, rules.min]"
            :type="show1 ? 'text' : 'password'"
            name="input-10-1"
            label="Password"
            hint="At least 3 characters"
            @click:append="show1 = !show1"
            class="mx-12 mb-5"
          ></v-text-field>
        </v-flex>
      </v-layout>
      <v-layout class="mt-0 pt-0">
        <v-flex class="d-flex justify-center">
          <v-btn
            @click="Loginvalidate"
            class="blue rounded-lg"
            :disabled="loginbtn"
          >
            LOG IN
          </v-btn>
        </v-flex>
      </v-layout>

      <p class="text-center mt-3">(or)</p>
      <!--   <hr class="mt-5 mb-1"> -->
      <v-layout>
        <v-flex class="d-flex justify-center flex-column">
          <!-- <v-btn @click="GetPasskey" class="red rounded-lg pa-5 white--text" :disabled="loginbtn"> -->
          <!-- <v-btn
            @click="GetPasskey"
            v-if="PasskeyState == 'I'"
            class="black rounded-lg pa-5 white--text caption"
          >
            Sign in with passkey
            <v-icon color="green">mdi-arrow-right-thick</v-icon>
          </v-btn> -->
          <v-btn @click="GetPasskey" class="red rounded-lg white--text">
            Add PassKey <v-icon>mdi-fingerprint</v-icon>
          </v-btn>
          <br /><br />
          <v-btn
            @click="StartAuthentication"
            v-if="PasskeyState == 'A'"
            class="red rounded-lg white--text"
          >
            Authenticate with Passkey <v-icon>mdi-fingerprint</v-icon>
          </v-btn>
        </v-flex>
      </v-layout>
    </v-card>
    <!-- {{ ENpasskey }}<br /><br />
    {{ DCpasskey }} -->
    <!-- {{ User_Details.userId }} -->
    <!-- {{ "Store" }} -->
    <!-- {{ this.$store.state.UserDetails }} -->
  </div>
</template>

<script>
import snackbar from "../snackbar.vue";
import EventService from "../../Services/EventService.js";

export default {
  name: "login",
  data() {
    return {
      // User_Details: {
      //   userId: "",
      //   password: "",
      //   displayName: "",
      // },
      ENpasskey: "",
      DCpasskey: "",
      User_Details: {
        userId: "",
        password: "",
        displayName: "",
      },
      data: "",
      obj: {
        snackbar: false,
        color: "red",
        topvar: false,
        bottomvar: false,
        text: "",
      },
      totaltime: "",
      show1: false,
      btnshow: false,

      rules: {
        required: (value) => !!value || "Required.",
        min: (v) => v.length >= 3 || "Min 3characters",
      },

      options: {},
      publicKeyCredentialCreationOptions: {},

      PasskeyState: "A",

      AuthenticationChallenge: {},
    };
  },
  components: {
    snackbar,
  },

  methods: {
    Loginvalidate() {
      if (this.User_Details.userId != "" && this.User_Details.password != "") {
        if (
          !(
            this.User_Details.userId[0] >= 0 || this.User_Details.userId[0] <= 9
          )
        ) {
          this.User_Details.userId = this.User_Details.userId.trim();
          this.User_Details.password = this.User_Details.password.trim();

          EventService.LoginValidationCall(this.User_Details)
            .then((response) => {
              if (response.data.status == "S") {
                this.$store.commit("setUserValue", response.data.userdetails);
                this.AssignCookies();

                this.obj = {
                  snackbar: true,
                  color: "green",
                  topvar: true,
                  bottomvar: false,
                  text: `${response.data.msg}`,
                };
                this.$router.push("/dashboard");
              } else if (response.data.status == "E") {
                this.obj = {
                  snackbar: true,
                  color: "red",
                  topvar: true,
                  bottomvar: false,
                  text: `${response.data.msg}`,
                };
              }
            })
            .catch((error) => {
              console.log(error);
            });
        } else {
          this.obj = {
            snackbar: true,
            color: "red",
            topvar: true,
            bottomvar: false,
            text: "User Id Does Not start with Numbers",
          };
        }
      } else {
        this.obj = {
          snackbar: true,
          color: "red",
          topvar: true,
          bottomvar: false,
          text: "Enter the Proper Data",
        };
      }
    },
    AssignCookies() {
      EventService.SetCookie(this.User_Details)
        .then((response) => {
          console.log(response);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    GetPasskey() {
      console.log("Im in passkey method");
      if (this.User_Details.userId != "") {
        this.GetDisplayName();
        EventService.ShowPasskeyRegistrationOptions(this.User_Details)
          .then((response) => {
            this.options = response.data;
            console.log("Response Data : ", response.data);

            this.publicKeyCredentialCreationOptions = {
              rp: {
                name: this.options.rp.name,
                id: this.options.rp.id,
              },
              user: {
                id: new TextEncoder().encode(this.options.user.id), // base64url-encoded user ID
                name: this.options.user.name,
                displayName: this.options.user.displayName,
              },
              challenge: this.base64urlToUint8Array(this.options.challenge),
              pubKeyCredParams: this.options.pubKeyCredParams,
              timeout: this.options.timeout || 60000, // default to 60 seconds if not provided
              excludeCredentials: this.options.excludeCredentials || [],
              authenticatorSelection: this.options.authenticatorSelection || {},
              attestation: this.options.attestation || "none",
              extensions: {
                // 🧼 Remove or conditionally include valid extensions only
                ...(this.options.extensions?.credProps && { credProps: true }),
                ...(this.options.extensions?.largeBlob && {
                  largeBlob: this.options.extensions.largeBlob,
                }),
                ...(this.options.extensions?.uvm && {
                  uvm: this.options.extensions.uvm,
                }),
                // Only add appidExclude if it's a valid non-null string
                ...(typeof this.options.extensions?.appidExclude === "string" &&
                  this.options.extensions.appidExclude.trim() !== "" && {
                    appidExclude: this.options.extensions.appidExclude,
                  }),
              },
            };

            console.log(
              "Options before save ",
              this.publicKeyCredentialCreationOptions
            );
            // ✅ Call it directly
            if ("credentials" in navigator) {
              console.log("WebAuthn is supported in this browser.");

              window.navigator.credentials
                .create({
                  publicKey: this.publicKeyCredentialCreationOptions,
                })
                .then((navResponse) => {
                  console.log("NAV Response : ", navResponse);

                  if (navResponse != null) {
                    const credential = navResponse;

                    const credentialData = {
                      id: credential.id,
                      rawId: this.bufferToBase64url(credential.rawId),
                      type: credential.type,
                      response: {
                        clientDataJSON: this.bufferToBase64url(
                          credential.response.clientDataJSON
                        ),
                        attestationObject: this.bufferToBase64url(
                          credential.response.attestationObject
                        ),
                      },
                      clientExtensionResults:
                        navResponse.getClientExtensionResults(),
                      authenticatorAttachment:
                        credential.authenticatorAttachment || null,
                    };

                    EventService.RegisterPasskey(
                      this.User_Details.userId,
                      credentialData
                    )
                      .then((response) => {
                        console.log("Response Data : ", response.data);
                        if (response.data.status == "S") {
                          this.obj = {
                            snackbar: true,
                            color: "green",
                            topvar: true,
                            bottomvar: false,
                            text: `${response.data.msg}`,
                          };
                          var UserDetails = {
                            Login_id: 1,
                            Role: "System Admin",
                            User_id: this.User_Details.userId,
                            History_id: 15,
                          };
                          this.$store.commit("setUserValue", UserDetails);
                          // this.$router.push("/dashboard");
                        } else if (response.data.status == "E") {
                          this.obj = {
                            snackbar: true,
                            color: "red",
                            topvar: true,
                            bottomvar: false,
                            text: response.data.errMsg,
                          };
                        }
                      })
                      .catch((error) => {
                        console.log(error);
                        this.obj = {
                          snackbar: true,
                          color: "red",
                          topvar: true,
                          bottomvar: false,
                          text: error.response,
                        };
                      });
                  }
                })
                .catch((navigatorerror) => {
                  console.log("ERROR : ", navigatorerror);
                });
            } else {
              console.log("WebAuthn is not supported in this browser.");
            }
            // if (response.data.status == "S") {
            //   this.ENpasskey = response.data.passkey;
            //   this.DCpasskey = this.decodeHex(this.ENpasskey);

            //   this.obj = {
            //     snackbar: true,
            //     color: "green",
            //     topvar: true,
            //     bottomvar: false,
            //     text: `${response.data.msg}`,
            //   };
            // } else if (response.data.status == "E") {
            //   this.obj = {
            //     snackbar: true,
            //     color: "red",
            //     topvar: true,
            //     bottomvar: false,
            //     text: response.data.errmsg,
            //   };
            // }
          })
          .catch((error) => {
            console.log(error);
            this.obj = {
              snackbar: true,
              color: "red",
              topvar: true,
              bottomvar: false,
              text: error.response,
            };
          });
      }
    },

    decodeHex(encodedString) {
      return Buffer.from(encodedString, "hex").toString("utf-8");
    },
    GetDisplayName() {
      this.User_Details.displayName =
        this.User_Details.userId.charAt(0).toUpperCase() +
        this.User_Details.userId.slice(1).toLowerCase();
    },
    base64urlToUint8Array(base64urlString) {
      const padding = "=".repeat((4 - (base64urlString.length % 4)) % 4);
      const base64 =
        base64urlString.replace(/-/g, "+").replace(/_/g, "/") + padding;
      const binary = window.atob(base64);
      return Uint8Array.from([...binary].map((c) => c.charCodeAt(0)));
    },
    bufferToBase64url(buffer) {
      return btoa(String.fromCharCode(...new Uint8Array(buffer)))
        .replace(/\+/g, "-")
        .replace(/\//g, "_")
        .replace(/=+$/, "");
    },

    TriggerNavigator() {},
    finishRegistration(navResponse) {
      const credential = navResponse;

      const credentialData = {
        id: credential.id,
        rawId: this.bufferToBase64url(credential.rawId),
        type: credential.type,
        response: {
          clientDataJSON: this.bufferToBase64url(
            credential.response.clientDataJSON
          ),
          attestationObject: this.bufferToBase64url(
            credential.response.attestationObject
          ),
        },
        clientExtensionResults: navResponse.getClientExtensionResults(),
        authenticatorAttachment: credential.authenticatorAttachment || null,
      };

      EventService.RegisterPasskey(this.User_Details.userId, credentialData)
        .then((response) => {
          console.log("Response Data : ", response.data);
          if (response.data.status == "S") {
            this.obj = {
              snackbar: true,
              color: "green",
              topvar: true,
              bottomvar: false,
              text: `${response.data.msg}`,
            };
          } else if (response.data.status == "E") {
            this.obj = {
              snackbar: true,
              color: "red",
              topvar: true,
              bottomvar: false,
              text: response.data.errMsg,
            };
          }
        })
        .catch((error) => {
          console.log(error);
          this.obj = {
            snackbar: true,
            color: "red",
            topvar: true,
            bottomvar: false,
            text: error.response,
          };
        });
    },
    StartAuthentication() {
      var User = {
        username: this.User_Details.userId,
      };
      // Step 1: Get authentication options
      EventService.startAuthentication(JSON.stringify(User))
        .then((response) => {
          console.log("Response Data : ", response.data);
          this.AuthenticationChallenge = response.data;
          console.log("Response Data : ", response.data);

          // Step 2: Decode challenge and allowCredentials
          // this.AuthenticationChallenge.challenge = Uint8Array.from(
          //   atob(this.AuthenticationChallenge.challenge),
          //   (c) => c.charCodeAt(0)
          // );

          this.AuthenticationChallenge.challenge = this.base64urlToUint8Array(
            this.AuthenticationChallenge.challenge
          );
          /* if (this.AuthenticationChallenge.allowCredentials != null) {
            this.AuthenticationChallenge.allowCredentials.forEach(
              (credential) => {
                credential.id = this.base64urlToUint8Array(credential.id);
              }
            );
          } */

          if (
            this.AuthenticationChallenge.allowCredentials &&
            Array.isArray(this.AuthenticationChallenge.allowCredentials)
          ) {
            this.AuthenticationChallenge.allowCredentials =
              this.AuthenticationChallenge.allowCredentials.map(
                (credential) => ({
                  ...credential,
                  id: this.base64urlToUint8Array(credential.id),
                })
              );
          }

          console.log(
            "Authentication Challenge : ",
            this.AuthenticationChallenge
          );

          // Step 3: Request credentials
          window.navigator.credentials
            .get({
              publicKey: this.AuthenticationChallenge,
            })
            .then((credential) => {
              console.log("Response Data : ", credential);
              // if (response.data.status == "S") {
              //   this.obj = {
              //     snackbar: true,
              //     color: "green",
              //     topvar: true,
              //     bottomvar: false,
              //     text: `${response.data.msg}`,
              //   };
              // const credential = response.data;

              // Step 4: Prepare data for backend
              const authData = {
                id: credential.id,
                rawId: this.bufferToBase64url(credential.rawId),
                type: credential.type,
                response: {
                  authenticatorData: this.bufferToBase64url(
                    credential.response.authenticatorData
                  ),

                  clientDataJSON: this.bufferToBase64url(
                    credential.response.clientDataJSON
                  ),

                  signature: this.bufferToBase64url(
                    credential.response.signature
                  ),
                  userHandle: credential.response.userHandle
                    ? this.bufferToBase64url(credential.response.userHandle)
                    : null,
                },
                clientExtensionResults: {},
              };

              // Step 5: Send to backend for verification
              EventService.finishAuthentication(authData)
                .then((response) => {
                  console.log("Response Data : ", response.data);
                })
                .catch((error) => {
                  console.log(error);
                });
              // console.log("Authenticated successfully", response.data);
              // this.$router.push("/dashboard");
              // } else if (response.data.status == "E") {
              //   this.obj = {
              //     snackbar: true,
              //     color: "red",
              //     topvar: true,
              //     bottomvar: false,
              //     text: response.data.errMsg,
              //   };
              // }
            })
            .catch((error) => {
              console.log(
                "Error in Navigator while fetching credentials from device" +
                  error
              );
              // this.obj = {
              //   snackbar: true,
              //   color: "red",
              //   topvar: true,
              //   bottomvar: false,
              //   text:
              //     "Error in Navigator while fetching credentials from device" +
              //     error.response,
              // };
            });
        })
        .catch((error) => {
          console.log(error);
          this.obj = {
            snackbar: true,
            color: "red",
            topvar: true,
            bottomvar: false,
            text: error.response,
          };
        });
    },
  },
  computed: {
    loginbtn() {
      if (this.User_Details.userId == "" || this.User_Details.password == "") {
        return true;
      } else {
        return false;
      }
    },
  },
};
</script>

<!-- // unmounted: {
      
    // } -->
<!-- // let currentdate = new Date().toJSON().slice(0, 10);
  // let hour = new Date().getHours() - 12;
  // let min = new Date().getMinutes();
  // let sec = new Date().getMilliseconds();
  // let time=new Date().toLocaleString();
  // console.log(time);
  // let totaltime = `${currentdate}-${hour}:${min}:${sec}`;
   // this.dt = `${currentdate}-${hour}:${min}:${sec}`;
  // alert(`${currentdate}-${hour}:${min}:${sec}`); -->

<!-- // this.$store.state.LoginHistory.Login_in=totaltime; -->

<!-- let currentdate= new Date().toJSON().slice(0,10);
  let hour=new Date().getHours();
  let min=new Date().getMinutes();
  let sec=new Date().getMilliseconds();
  
  
  let totaltime=`${currentdate}-${hour}:${min}:${sec}`;
  
  // this.array1.push({UserId:this.userid,Login_in:totaltime,Log_out:this.logout});
  
  this.$store.state.LoginHistory.push({UserId:this.userid,Login_in:totaltime,Log_out:'null'});
  this.dt=`${currentdate}-${hour}:${min}:${sec}`;
  alert(`${currentdate}-${hour}:${min}:${sec}`); -->

<!-- // arr.forEach(elem => {
    //   console.log(elem);

    //   if (elem.UserId == this.userid && elem.Password == this.password) {
    //     console.log('jdfjhj');
    //     alert('successfully Logined');
    //     this.prole=elem.Role;
    //     this.datetimelog();
    //     alert('successfully Inserted');
    //     this.$router.push('/dashboard');
    //     break;
        
    //   } else {
    //     alert('Not successfully logined');
    //     break;
    //   }

    // }); -->
