<template>

    <div>
      <v-card class="blue lighten-3 d-flex justify-center align-center" :fluid="true" style="height: 100vh;">
      <!-- <loginpage v-show="this.$store.state.UserDetails.Role !='' ? false : true" style="width: 450px;height: 500px;" /> -->
      <loginpage  v-show="this.$store.state.UserDetails.Role !='' ? false : true"  style="width: 450px;height: 500px;" />
    </v-card>
    </div>
    
   
  
</template>
<script>
// import EventService from '../Services/EventService';

import loginpage from '../components/login/loginpage.vue';
export default {
  name: 'login',
  components: {
    loginpage,
  },
  // beforeMount(){
  //   EventService.GetCookie()
  //   .then((response)=>{
  //     console.log(response);
  //   })
  //   .catch((error)=>{
  //     console.log(error);
  //   })
  // }


}
</script>



 <!--  <snackbar v-show="val" /> 
    <v-container>

      <v-card class="d-flex flex-column mt-6 grey lighten-3 pa-6 mx-auto" max-width="500">
        <v-row d-flex justify="center">
          <h1 class="mb-4">Login Page</h1>
        </v-row>
        <v-row d-flex justify="center">
          <v-col cols="12" sm="8">
            <v-text-field v-model="userid" label="UserId" outlined class="pl-5" type="text"></v-text-field>
          </v-col>
        </v-row>

        <v-row d-flex justify="center">
          <v-col cols="12" sm="8">
            <v-text-field v-model="password" label="Password" outlined class="pl-5" type="password"></v-text-field>
          </v-col>
        </v-row>

        <v-row d-flex justify="center">
          <v-btn color="blue" @click=" validate1" class="mb-6">
            LOG IN
          </v-btn>
        </v-row>

      </v-card>
      {{ typeof(new Date().toJSON().stringify()) }}<br>
      {{'jh'+ prole }}<br>
      {{ password }}<br>
      {{ $store.state.LoginHistory }}<br>
      {{ $store.state.Login }}<br>
    </v-container> -->

    <!-- // data() {
      //   return {
      //     userid:'',
      //     password:'',
      //     dt:'',
      //     val:false,
      //     prole:'',
      //   }
      // },
      // methods: {
      //   validate1() {
      //     let arr = this.$store.state.Login;
      //     console.log(arr);
      //     if (this.userid !='' && this.password !='') {
      //       for(let i=0;i<arr.length;i++){
      //         if(arr[i].UserId==this.userid && arr[i].Password==this.password){
      //           alert('successfully Logined');
      //           this.prole=arr[i].Role;
      //           console.log(this.prole);
      //           this.datetimelog();
      //           alert('successfully Inserted');
      //           this.$router.push('/dashboard');
      //           break;
      //         }else{
      //           alert('value not matched');
      //         }
      //       }        
      //     }
      //     else {
      //       console.log('enter the proper data');
      //       this.val = true;
      //     }
      //   },
      //   datetimelog() {
      //     let totaltime=new Date().toLocaleString();
      //     this.$store.state.LoginHistory.push({ UserId: this.userid, Login_in: totaltime, Log_out: 'null' });
      //   }
      // }, -->
      




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
  //     alert('suthis.userrole=data;ccessfully Inserted');
  //     this.$router.push('/dashboard');
  //     break;
      
  //   } else {
  //     alert('Not successfully logined');
  //     break;
  //   }

  // }); -->