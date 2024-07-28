<template>
  <div class="login">
    <snackbar :obj="obj" />

    <v-card class="grey pa-6 lighten-4 mx-auto rounded-xl" width="400" height="450">
      <v-layout>
        <v-flex>
          <h1 class="text-center mb-8 font-weight-black grey--text text-capitalize text--darken-2">
            Login Page
          </h1>
        </v-flex>
      </v-layout>
      <v-layout>
        <v-flex>
          <v-text-field v-model="User_Details.user_id" label="UserId" outlined type="text"
            class="mx-12 mb-5"></v-text-field>
        </v-flex>
      </v-layout>
      <v-layout>
        <v-flex><v-text-field outlined v-model="User_Details.password" :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
            :rules="[rules.required, rules.min]" :type="show1 ? 'text' : 'password'" name="input-10-1" label="Password"
            hint="At least 3 characters" @click:append="show1 = !show1" class="mx-12 mb-5"></v-text-field>
        </v-flex>
      </v-layout>
      <v-layout>
        <v-flex class="d-flex justify-center">
          <v-btn @click="Loginvalidate" class="blue rounded-lg pa-5" :disabled="loginbtn">
            LOG IN
          </v-btn>
        </v-flex>
      </v-layout>
      <v-layout>
        <v-flex class="d-flex justify-center mt-5">
          <v-btn @click="GetPasskey" class="red rounded-lg pa-5 white--text" :disabled="loginbtn">
            Create PassKey <v-icon>mdi-fingerprint</v-icon>
          </v-btn>
        </v-flex>
      </v-layout>
    </v-card>
    {{ ENpasskey }}<br><br>
    {{ DCpasskey }}
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
      ENpasskey: "",
      DCpasskey: "",
      User_Details: {
        user_id: "",
        password: "",
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
    };
  },
  components: {
    snackbar,
  },

  methods: {
    Loginvalidate() {
      if (this.User_Details.user_id != "" && this.User_Details.password != "") {
        if (
          !(
            this.User_Details.user_id[0] >= 0 ||
            this.User_Details.user_id[0] <= 9
          )
        ) {
          this.User_Details.user_id = this.User_Details.user_id.trim();
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
      EventService.PasskeyGen(this.User_Details.user_id)
        .then((response) => {
          console.log(response.data);
          if (response.data.status == "S") {
            this.ENpasskey = response.data.passkey;
            this.DCpasskey = this.decodeHex(this.ENpasskey)

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
              text: response.data.errmsg,
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
    decodeHex(encodedString) {
      return Buffer.from(encodedString, "hex").toString("utf-8");
    },
  },
  computed: {
    loginbtn() {
      if (this.User_Details.user_id == "" || this.User_Details.password == "") {
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
