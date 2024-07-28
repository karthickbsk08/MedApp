 <template>
  <div>

    <v-layout></v-layout>
    <v-layout class="mt-16"></v-layout>
    <v-layout>
      <v-container>
        <v-layout justify="center">
          <v-expansion-panels color="black" flat class="pa-5">
            <v-expansion-panel class="mt-16 rounded-t-lg">
              <v-expansion-panel-header
                class="blue lighten-3 font-weight-black grey--text text--darken-3 text-body-1 text-end " :open="true">ADD
                USER </v-expansion-panel-header>
              <v-expansion-panel-content class="grey lighten-2 rounded-b-xl">
                <v-layout class="mt-5 mx-auto ">
                  <v-col cols="12" sm="3" md="3">
                    <v-text-field v-model="NewUser.user_id" label="User Id" outlined class="pl-2 rounded-lg" type="text"
                      max-width="100" dense></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="3" md="3">
                    <v-text-field v-model="NewUser.password" label="Password" outlined class="pl-2 rounded-lg"
                      type="password" max-width="100" dense></v-text-field>
                    <!-- <v-text-field outlined v-model="password" :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
                      :rules="[rules.required, rules.min]" :type="show1 ? 'text' : 'password'" name="input-10-1"
                      label="Password" hint="At least 3 characters" @click:append="show1 = !show1"
                      class="pl-2 rounded-lg" dense ></v-text-field> -->
                  </v-col>
                  <v-col cols="12" sm="3" md="3">
                    <v-select :items="items" label="Role" outlined max-width="100" dense v-model="NewUser.role"
                      class="pl-2"></v-select>
                  </v-col>
                  <v-col class="d-flex justify-center pl-2" sm="3" md="3">
                    <v-btn color="grey darken-2 white--text px-5 rounded-xl" @click="adduser">ADD</v-btn>
                  </v-col>
                </v-layout>
              </v-expansion-panel-content>
            </v-expansion-panel>
          </v-expansion-panels>

          <snackbar :obj="obj" />
          <br><br>
          <!-- {{ Resulttemp }}<br><br> 
           {{ this.$store.state.Login }}  -->

        </v-layout>
      </v-container>



    </v-layout>


  </div>
</template> 
  
<script>

import EventService from '../Services/EventService.js'
import snackbar from '../components/snackbar.vue';

export default {
  name: "Adduser",
  components: {
    snackbar,
  },
  data() {
    return {
      items: ['System Admin', 'Biller', 'Manager', 'Inventory'],
      NewUser: {
        user_id: "",
        password: "",
        role: "",
        cuser_id: 0
      },
      Resulttemp: "",
      obj: {
        snackbar: false,
        color: 'red',
        topvar: false,
        bottomvar: false,
        location: "center",
        text: '',
      },
      rules: {
        required: value => !!value || 'Required.',
        min: v => v.length >= 3 || 'Min 3characters',
      }
    }
  },
  methods: {
    adduser() {
      this.NewUser.user_id = this.NewUser.user_id.trim();
      this.NewUser.password = this.NewUser.password.trim();
      this.NewUser.cuser_id = this.$store.state.UserDetails.User_id
      console.log(this.NewUser.cuser_id);

      if (this.NewUser.user_id != '' && this.NewUser.password != '' && this.NewUser.role != '') {
        EventService.AddNewUser(this.NewUser)
          .then((response) => {
            console.log(response);
            this.Resulttemp = response.data

            if (response.data.status == "S") {

              this.obj = {
                snackbar: true,
                color: 'green',
                topvar: true,
                bottomvar: false,
                text: `Welcome to ${this.NewUser.user_id.toUpperCase()}`,
              };
              this.NewUser.user_id = ""
              this.NewUser.password = ""
              this.NewUser.role = ""
            } else if (response.data.status == "E") {

              this.obj = {
                snackbar: true,
                color: 'red',
                topvar: true,
                bottomvar: false,
                location: "right",
                text: `${response.data.msg}`,
              };
              this.NewUser.user_id = ""
              this.NewUser.password = ""
              this.NewUser.role = ""
            }
          })
          .catch((error) => {
            console.log(error);
          });
      } else {
        this.obj = {
          snackbar: true,
          color: 'red',
          topvar: true,
          bottomvar: false,
          location: "right",
          text: 'Enter the Valid Inputs',
        };
        this.NewUser.user_id = ""
        this.NewUser.password = ""
        this.NewUser.role = ""
      }

    }
  }
}





</script>


