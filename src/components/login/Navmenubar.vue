<template>
    <div class="Navmenubar">
        <v-navigation-drawer v-model="drawer" app class=" grey lighten-2" >

            <v-list-item class="text-h6 d-flex flex-column grey lighten-2 ">
                <v-list-item-avatar height="150" width="150"
                    class="blue lighten-1 text-h1 mt-5 justify-center white--text">{{ UserName.slice(0, 1).toUpperCase() }}
                </v-list-item-avatar>
                <v-list-item-content>

                    <v-list-item-title class="grey--text text--darken-1 text-h5 font-weight-bold text-left">{{UserName.toUpperCase()
                    }}
                        <p class="pt-2 green--text text--darken-3 text-left pb-0">{{ this.$store.state.UserDetails.Role }}
                        </p>
                    </v-list-item-title>
                </v-list-item-content>
            </v-list-item>
            <v-divider></v-divider>
        </v-navigation-drawer>

        <v-app-bar app dark class="blue lighten-2">

            <v-avatar size="60" @click="drawer = !drawer">
                <v-icon size="40">
                    mdi-account-circle
                </v-icon>
            </v-avatar>

            <v-toolbar-title class="white--text font-weight-bold text-h5 text-uppercase"><span
                    class="grey--text text--darken-1">Med</span>-App</v-toolbar-title>
            <v-spacer></v-spacer>

            <div class="d-none d-md-flex">
                <v-btn text color="grey darken-3" class="font-weight-black" to="/dashboard" v-show="show1">
                    <v-icon left>mdi-view-dashboard</v-icon>
                    <span>DASHBOARD</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/adduser" v-show="show2">
                    <v-icon left>mdi-account</v-icon>
                    <span>ADD USER</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/loginhistory" v-show="show3">
                    <v-icon left>mdi-history</v-icon>
                    <span>LOGINHISTORY</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/stockentry" v-show="show4">
                    <v-icon left>mdi-basket-fill</v-icon>
                    <span>STOCK ENTRY</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/stockview" v-show="show5">
                    <v-icon left>mdi-teamviewer</v-icon>
                    <span>STOCK VIEW</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/billentry" v-show="show6">
                    <v-icon left>mdi-receipt</v-icon>
                    <span>BILLENTRY</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/salesreport" v-show="show7">
                    <v-icon left>mdi-sale</v-icon>
                    <span>SALESREPORT</span>
                </v-btn>
            </div>

            <v-dialog v-model="dialog" max-width="350" class="pa-6">
                <template v-slot:activator="{ on, attrs }" v-show="show8">
                    <v-btn class="white--text darken-1 grey font-weight-black ml-5" dark v-bind="attrs" v-on="on">
                        LOGOUT
                    </v-btn>
                </template>
                <v-card class="rounded-lg">
                    <v-card-title class="text-h6  text-center text-uppercase blue lighten-1 pa-8 white--text">
                        Do You Want To Logout
                    </v-card-title>

                    <v-card-actions class="pa-5">
                        <v-spacer></v-spacer>
                        <v-btn color="red darken-1" text @click="dialog = false">
                            Disagree
                        </v-btn>
                        <v-btn color="green darken-1" text @click="DoLogout">
                            Agree
                        </v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
            <mobiledrawer :Role="role" class="d-flex d-md-none ml-2" />
        </v-app-bar>
    </div>
</template>
<script>
import Eventservice from '../../Services/EventService.js'
import mobiledrawer from './Mobiledrawer.vue';
export default {
    name: 'App',
    components: {
        mobiledrawer,
    },
    watch: {
        role() {
            if (this.role == 'System Admin') {
                this.show1 = true;
                this.show2 = true;
                this.show3 = true;
                this.show4 = false;
                this.show5 = false;
                this.show6 = false;
                this.show7 = false;

            } else if (this.role == "Inventry") {
                this.show1 = true;
                this.show2 = false;
                this.show3 = false;
                this.show4 = true;
                this.show5 = true;
                this.show6 = false;
                this.show7 = false;

            } else if (this.role == "Biller") {
                this.show1 = true;
                this.show2 = false;
                this.show3 = false;
                this.show4 = false;
                this.show5 = true;
                this.show6 = true;
                this.show7 = false;

            } else if (this.role == "Manager") {
                this.show1 = true;
                this.show2 = false;
                this.show3 = false;
                this.show4 = true;
                this.show5 = true;
                this.show6 = false;
                this.show7 = true;

            }
        }
    },
    updated() {
        this.UserName = this.$store.state.UserDetails.User_id
    },
    data() {
        return {
            drawer: null,
            show1: false,
            show2: false,
            show3: false,
            show4: false,
            show5: false,
            show6: false,
            show7: false,
            show8: true,
            // role:'',  
            UserName: this.$store.state.UserDetails.User_id,
            items: [
                { title: 'DASHBOARD', icon: 'mdi-view-dashboard', route: '/dashboard' },
                { title: 'LOGIN HISTORY', icon: 'mdi-login', route: '/loginhistory' },
                { title: 'ADD USER', icon: 'mdi-account', route: '/adduser' },
                { title: 'STOCK VIEW', icon: 'mdi-teamviewer', route: '/stockview' },
                { title: 'STOCK ENTRY', icon: 'mdi-basket-fill', route: '/stockentry' },
                { title: 'BILL ENTRY', icon: 'mdi-receipt', route: '/billentry' },
                { title: 'SALES REPORT', icon: 'mdi-sale', route: '/salesreport' },
            ],

            logouttime: '',
            logintime: '',
            dialog: false,
            LogoutDetails: {
                login_history_id: 0,
                user_id: "",
            }


        }
    },
    props: {
        role: String
    },
    methods: {

        DoLogout() {
            this.LogoutDetails.login_history_id = this.$store.state.UserDetails.History_id
            this.LogoutDetails.user_id = this.$store.state.UserDetails.User_id
            // console.log(this.LogoutDetails);
            // this.LogoutDetails = JSON.stringify(this.LogoutDetails)

            Eventservice.Logout(JSON.stringify(this.LogoutDetails))
                .then((response) => {
                    console.log(response);
                    if (response.data.status == "S") {
                        // console.log(response.data.msg);
                        this.dialog = false;
                        this.$store.commit('SetEmptyUserDetails')
                        this.$router.push('/');

                    } else if (response.data.status == "E") {
                        console.log(response.data.msg);
                    }
                })
                .catch((error) => {
                    console.log(error);
                })
        }
        // datetimelog() {

        //     this.logouttime = new Date().toLocaleString();
        //     // console.log(this.logouttime);
        //     for (let i = 0; i < this.LoginHistoryarr.length; i++) {
        //         if (this.LoginHistoryarr[i].Log_out == 'In_Working') {

        //             // alert(this.logouttime);
        //             this.LoginHistoryarr[i].Log_out = this.logouttime;
        //             // this.role='';
        //             // console.log(this.role);
        //             // this.$emit('logouttime');
        //             this.$store.state.Temp.role = ""
        //             // this.$emit('nav');
        //             this.dialog = false
        //             this.$router.push('/');

        //             break;
        //         }
        //     }

        //     // let currentdate = new Date().toJSON().slice(0, 10);
        //     // let hour = new Date().getHours() - 12;
        //     // let min = new Date().getMinutes();
        //     // let sec = new Date().getMilliseconds().slice(0,2);
        //     // // let totaltime = `${currentdate}-${hour}:${min}:${sec}`;
        //     // this.logouttime = `${currentdate}-${hour}:${min}:${sec}`;

        // },

    },
    // mounted() {
    //     this.role=this.$store.state.UserDetails.Role
    //     console.log("NAv bar",this.role); 


    // }


}
</script>

<!--                                        TOOL BAR -->
  <!-- <v-navigation-drawer v-model="drawer" >
                
            </v-navigation-drawer>  -->

 <!-- <v-toolbar dark class="blue lighten-2">
                <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
                <v-toolbar-title class="white--text font-weight-bold">Med-App</v-toolbar-title>
                <v-spacer></v-spacer>
                <v-btn text color="black" class="font-weight-black">
                    <v-icon left>mdi-view-dashboard</v-icon>
                    <span>DASHBOARD</span>
                </v-btn>
                <v-btn text color="black" class="font-weight-black">
                    <v-icon left>mdi-account</v-icon>
                    <span>ADD USER</span>
                </v-btn>
                <v-btn text color="black" class="font-weight-black">
                    <v-icon left>mdi-history</v-icon>
                    <span>LOGIN HISTORY</span>
                </v-btn>
                <v-btn text color="black" class="font-weight-black">
                    <v-icon left>mdi-basket-fill</v-icon>
                    <span>STOCK ENTRY</span>
                </v-btn>
                <v-btn text color="black" class="font-weight-black">
                    <v-icon left>mdi-teamviewer</v-icon>
                    <span>STOCK VIEW</span>
                </v-btn>
                <v-btn text color="black" class="font-weight-black">
                    <v-icon left>mdi-receipt</v-icon>
                    <span>BILL ENTRY</span>
                </v-btn>
                <v-btn text color="black" class="font-weight-blgstack">
                    <v-icon left>mdi-sale</v-icon>
                    <span>SALES REPORT</span>
                </v-btn>
                <v-btn text color="black" class="font-weight-black" @click="datetimelog">
                    <v-icon left>mdi-logout-variant</v-icon>
                    <span>LOG OUT</span>
                </v-btn>


            </v-toolbar>-->

             <!-- </v-card> -->
            <!-- <v-navigation-drawer v-model="drawer" app>
                <v-list-item>
                    <v-list-item-content>
                        <v-list-item-title class="text-h6">
                            Application
                        </v-list-item-title>
                        <v-list-item-subtitle>
                            subtext
                        </v-list-item-subtitle>
                    </v-list-item-content>
                </v-list-item>

                <v-divider></v-divider>

                <v-list dense nav>
                    <v-list-item v-for="item in items" :key="item.title" :to="item.route">
                        <v-list-item-icon>  
                            <v-icon>{{ item.icon }}</v-icon>
                        </v-list-item-icon>

                        <v-list-item-content>
                            <v-list-item-title>{{ item.title }}</v-list-item-title>
                        </v-list-item-content>
                    </v-list-item>
                </v-list>
            </v-navigation-drawer>  <div v-if="role == 'System Admin'">
                <v-btn text color="grey darken-3" class="font-weight-black" to="/dashboard">
                    <v-icon left>mdi-view-dashboard</v-icon>
                    <span>DASHBOARD</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/adduser">
                    <v-icon left>mdi-account</v-icon>
                    <span>ADD USER</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/loginhistory">
                    <v-icon left>mdi-history</v-icon>
                    <span>LOGINHISTORY</span>
                </v-btn>
            </div>

            <div v-else-if="role == 'Inventory'">
                <v-btn text color="grey darken-3" class="font-weight-black" to="/stockentry">
                    <v-icon left>mdi-basket-fill</v-icon>
                    <span>STOCK ENTRY</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/stockview">
                    <v-icon left>mdi-teamviewer</v-icon>
                    <span>STOCK VIEW</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/dashboard">
                    <v-icon left>mdi-view-dashboard</v-icon>
                    <span>DASHBOARD</span>
                </v-btn>
            </div>

            <div v-else-if="role == 'Biller'">

                <v-btn text color="grey darken-3" class="font-weight-black" to="/billentry">
                    <v-icon left>mdi-receipt</v-icon>
                    <span>BILLENTRY</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/stockview">
                    <v-icon left>mdi-teamviewer</v-icon>
                    <span>STOCK VIEW</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/dashboard">
                    <v-icon left>mdi-view-dashboard</v-icon>
                    <span>DASHBOARD</span>
                </v-btn>

            </div>

            <div v-else-if="role == 'Manager'">
                <v-btn text color="grey darken-3" class="font-weight-black" to="/salesreport">
                    <v-icon left>mdi-sale</v-icon>
                    <span>SALESREPORT</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/stockview">
                    <v-icon left>mdi-teamviewer</v-icon>
                    <span>STOCKVIEW</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/stockentry">
                    <v-icon left>mdi-basket-fill</v-icon>
                    <span>STOCKENTRY</span>
                </v-btn>
                <v-btn text color="grey darken-3" class="font-weight-black" to="/dashboard">
                    <v-icon left>mdi-view-dashboard</v-icon>
                    <span>DASHBOARD</span>
                </v-btn>
            </div> -->

            <!-- <v-btn text color="black" class="font-weight-black" @click="datetimelog">
                <v-icon left>mdi-logout-variant</v-icon>
                <span>LOG OUT</span>
            </v-btn> -->

