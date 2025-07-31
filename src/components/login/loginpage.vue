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
      console.log("Im in passkey INPUT : ", JSON.stringify(this.User_Details));
      if (this.User_Details.userId != "") {
        this.GetDisplayName();
        EventService.ShowPasskeyRegistrationOptions(this.User_Details)
          .then((response) => {
            this.options = response.data;
            console.log("Response Data : ", JSON.stringify(response.data));

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

            // console.log(
            //   "Options before save ",
            //   this.publicKeyCredentialCreationOptions
            // );
            // ✅ Call it directly
            if ("credentials" in navigator) {
              console.log("WebAuthn is supported in this browser.");

              console.log(
                "PublicKeyCredentialCreationOptions : ",
                JSON.stringify(this.publicKeyCredentialCreationOptions)
              );

              var dummy = {
                publicKey: this.publicKeyCredentialCreationOptions,
              };

              console.log("Navigator create() ==>  : ", JSON.stringify(dummy));
              window.navigator.credentials
                .create(dummy)
                .then((navResponse) => {
                  console.log(
                    "NAV Response ==> : ",
                    JSON.stringify(navResponse)
                  );

                  if (navResponse != null) {
                    const credential = navResponse;

                    // const credentialData = {
                    //   id: credential.id,
                    //   rawId: this.bufferToBase64url(credential.rawId),
                    //   type: credential.type,
                    //   response: {
                    //     clientDataJSON: this.bufferToBase64url(
                    //       credential.response.clientDataJSON
                    //     ),
                    //     attestationObject: this.bufferToBase64url(
                    //       credential.response.attestationObject
                    //     ),
                    //   },
                    //   clientExtensionResults:
                    //     navResponse.getClientExtensionResults(),
                    //   authenticatorAttachment:
                    //     credential.authenticatorAttachment || null,
                    // };
                    console.log(
                      "Register Passkey input body: ",
                      JSON.stringify(credential)
                    );

                    EventService.RegisterPasskey(
                      this.User_Details.userId,
                      credential
                    )
                      .then((response) => {
                        console.log("Response Data : ", response.data);
                        console.log(
                          "Register Passkey input body: ",
                          JSON.stringify(response.data)
                        );

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

      {
        /*    "challenge": {
        "0": 210,
        "1": 100,
        "2": 108,
        "3": 159,
        "4": 8,
        "5": 110,
        "6": 117,
        "7": 10,
        "8": 93,
        "9": 16,
        "10": 107,
        "11": 5,
        "12": 28,
        "13": 201,
        "14": 247,
        "15": 14,
        "16": 97,
        "17": 139,
        "18": 215,
        "19": 146,
        "20": 231,
        "21": 104,
        "22": 93,
        "23": 164,
        "24": 198,
        "25": 231,
        "26": 65,
        "27": 54,
        "28": 159,
        "29": 10,
        "30": 163,
        "31": 71
    },
    "timeout": null,
    "rpId": "localhost",
    "allowCredentials": [],
    "userVerification": "preferred",
    "extensions": {
        "appid": null,
        "largeBlob": null,
        "uvm": null
    } */
      }
      // Step 1: Get authentication options
      EventService.startAuthentication(JSON.stringify(User))
        .then((response) => {
          console.log("Response Data : ", response.data);
          this.AuthenticationChallenge = response.data;
          console.log("Response Data : ", response.data);

          console.log("Authentication Challenge : ", {
            ...this.AuthenticationChallenge,
          });
          /*{
    "challenge": "szd0bOkOTFcQXdTTHTgsuPWTUBh2l-Luv2mr6rvpuIE",
    "timeout": null,
    "rpId": "localhost",
    "allowCredentials": [
        {
            "type": "public-key",
            "id": "Eti28oNsm68oNJQxIR_wYg",
            "transports": null
        }
    ],
    "userVerification": "preferred",
    "extensions": {
        "appid": null,
        "largeBlob": null,
        "uvm": null
    }
}*/
          console.log(
            "AuthenticationChallenge Challenge : ",
            Object.values(
              this.AuthenticationChallenge.publicKeyCredentialRequestOptions
                .challenge
            )
          );

          console.log(
            "AuthenticationChallenge bytearray : ",
            new Uint8Array(
              Object.values(
                this.AuthenticationChallenge.publicKeyCredentialRequestOptions
                  .challenge
              )
            )
          );

          console.log(
            "AuthenticationChallenge bytearray : ",
            new Uint8Array(
              Object.values(
                this.AuthenticationChallenge.publicKeyCredentialRequestOptions
                  .challenge
              )
            ).buffer
          );

          console.log(
            "this.AuthenticationChallenge.allowCredentials : ",
            this.AuthenticationChallenge.publicKeyCredentialRequestOptions
              .allowCredentials
          );

          /*       {
    "challenge": {},
    "timeout": null,
    "rpId": "localhost",
    "allowCredentials": [
        {
            "type": "public-key",
            "id": {}
        }
    ],
    "userVerification": "preferred",
    "extensions": {
        "appid": null,
        "largeBlob": null,
        "uvm": null
    }
} */

          /*   const publicKey = {
            ...this.AuthenticationChallenge,
            challenge: this.base64urlToUint8Array(
              this.AuthenticationChallenge.challenge
            ).buffer,
            allowCredentials: this.AuthenticationChallenge.allowCredentials.map(
              (cred) => {
                console.log("Credential : ", { ...cred });
                const cleaned = {
                  ...cred,
                  id: this.base64urlToUint8Array(cred.id).buffer,
                };
                // Remove 'transports' if null or undefined
                if (!Array.isArray(cred.transports)) {
                  delete cleaned.transports;
                }
                return cleaned;
              }
            ),
          };

          console.log("Public Key : ", publicKey);
 */

          console.log(" ...this.AuthenticationChallenge : ", {
            ...this.AuthenticationChallenge,
          });

          /*       
           {
    "publicKeyCredentialRequestOptions": {
        "challenge": "BK5ZfwbajEplu6LfOSt4WbTlPuT7G38UQ-zTwr1gMeY",
        "timeout": null,
        "rpId": "localhost",
        "allowCredentials": [
            {
                "type": "public-key",
                "id": "G5Vspl6izBEk66d6QRCEcg",
                "transports": null
            }
        ],
        "userVerification": "preferred",
        "extensions": {
            "appid": null,
            "largeBlob": null,
            "uvm": null
        }
    },
    "username": null,
    "userHandle": "td86BIAlw7-INzPeJWBItMWgawRHSnWrnD6ITN18ZPA"
}
} */
          // const { extensions: _, ...cleanedChallenge } =
          //   this.AuthenticationChallenge;

          const publicKey = {
            ...this.AuthenticationChallenge.publicKeyCredentialRequestOptions,
            extensions: {},
            challenge: this.base64urlToUint8Array(
              this.AuthenticationChallenge.publicKeyCredentialRequestOptions
                .challenge
            ).buffer,
            allowCredentials:
              this.AuthenticationChallenge.publicKeyCredentialRequestOptions.allowCredentials.map(
                (cred) => {
                  const formatted = {
                    ...cred,
                    id: this.base64urlToUint8Array(cred.id).buffer,
                  };
                  if (!Array.isArray(cred.transports)) {
                    delete formatted.transports;
                  }
                  return formatted;
                }
              ),
            username: "",
            timeout: "",
            userHandle: this.AuthenticationChallenge.userHandle
              ? this.base64urlToUint8Array(
                  this.AuthenticationChallenge.userHandle
                ).buffer
              : null,
          };

          console.log("publicKey : ", publicKey);

          // Step 2: Decode challenge and allowCredentials
          // this.AuthenticationChallenge.challenge = Uint8Array.from(
          //   atob(this.AuthenticationChallenge.challenge),
          //   (c) => c.charCodeAt(0)
          // );

          // this.AuthenticationChallenge.challenge = this.base64urlToUint8Array(
          //   this.AuthenticationChallenge.challenge
          // );
          /* if (this.AuthenticationChallenge.allowCredentials != null) {
            this.AuthenticationChallenge.allowCredentials.forEach(
              (credential) => {
                credential.id = this.base64urlToUint8Array(credential.id);
              }
            );
          } */

          /*          {
    "challenge": "6KZF2wfd5pjeqS64UjdEmDNxvMWTykS4P3Wj3aVjuLY",
    "timeout": null,
    "rpId": "localhost",
    "allowCredentials": [
        {
            "type": "public-key",
            "id": "4cVenuCyO3U481dOecvHtA",
            "transports": null
        }
    ],
    "userVerification": "preferred",
    "extensions": {
        "appid": null,
        "largeBlob": null,
        "uvm": null
    }
        // Step 2: Decode base64url fields
      const publicKey = {
        ...options,
        challenge: this.base64urlToBuffer(options.challenge),
        allowCredentials: options.allowCredentials.map(cred => ({
          ...cred,
          id: this.base64urlToBuffer(cred.id)
        }))
      };

      // Step 3: Call WebAuthn API
      const assertion = await navigator.credentials.get({ publicKey });

      console.log("Assertion result:", assertion);
} */

          /*    if (
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
 */

          // console.log(
          //   "challenge:",
          //   this.bufferToBase64url(publicKey.challenge)
          // );
          // console.log(
          //   "credential ID:",
          //   this.bufferToBase64url(publicKey.allowCredentials[0].id)
          // );

          // console.log("{ ...publicKey } : ", { ...publicKey });

          // Step 3: Request credentials
          navigator.credentials
            .get({ publicKey })
            .then((credential) => {
              console.log("GET CREDENTIAL : ", credential);
              console.log(
                "GET CREDENTIAL stringify : ",
                JSON.stringify(credential)
              );

              /*          {
    "authenticatorAttachment": "platform",
    "clientExtensionResults": {},
    "id": "0GiwWfVX1RVa6khv-zDsOg",
    "rawId": "0GiwWfVX1RVa6khv-zDsOg",
    "response": {
        "authenticatorData": "SZYN5YgOjGh0NBcPZHZgW4_krrmihjLHmVzzuoMdl2MZAAAAAA",
        "clientDataJSON": "eyJ0eXBlIjoid2ViYXV0aG4uZ2V0IiwiY2hhbGxlbmdlIjoiRzV3aVhnSjVqQ09mUXhxX1ZKRzN5ODBfWTNCcVZsWHYxc1NrN09VR3FERSIsIm9yaWdpbiI6Imh0dHA6Ly9sb2NhbGhvc3Q6ODA4MCIsImNyb3NzT3JpZ2luIjpmYWxzZX0",
        "signature": "MEUCIQDBnB7GalY_Z2TAXBM5AHqa8hzrPoxJMVi8GKyTe7YmvQIgHD1IM3UVBH0QDk1O6EHs6GM0CuTPgclprPXK5lWI814",
        "userHandle": "eTJuRlpxeFdlbzV1MmdCa3dRdVFSUTlzYmZDSG9Fd2FQNGxka252X29OTQ"
    },
    "type": "public-key"
}
   */
              const authData = {
                id: credential.id,
                rawId: this.bufferToBase64url(credential.rawId),
                type: credential.type,
                authenticatorAttachment: credential.authenticatorAttachment,
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
                clientExtensionResults:
                  credential.getClientExtensionResults?.() || {},
              };

              return EventService.finishAuthentication(authData)
                .then((res) => {
                  console.log("finishAuthentication : " + res);
                })
                .catch((err) => {
                  console.log("finishAuthentication : " + err);
                });
            })
            // .then((response) => {
            //   console.log("Authentication success:", response.data);
            //   // Redirect or show success message
            // })
            .catch((error) => {
              console.log("WebAuthn error stringify : ", JSON.stringify(error));
              console.log("WebAuthn error:", error);
              // Show error message
            });
          /*  window.navigator.credentials
            .get({
              publicKey: publicKey,
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
            }); */
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
