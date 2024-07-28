<template>
  <div class="mt-16">
    <div v-if="role == 'Biller'">
      <Billerdash  />
    </div>
    <div v-else-if="role == 'System Admin'">
      <SystemAdmin />
    </div>
    <div v-else-if="role == 'Manager'">
      <Manager />
    </div>
    <div v-else-if="role == 'Inventry'">
      <SystemAdmin />
    </div>  
  </div>
</template>

<script >

import SystemAdmin from '../components/dashboard/SystemAdmin.vue';
import Billerdash from '../components/dashboard/Billerdash.vue';
import Manager from '../components/dashboard/Manager/Managerdash.vue';


export default {
  name: "dashboard",
  components: {
    SystemAdmin,
    Billerdash,
    Manager,


  },
  data() {
    return {
      cal: false,
      role: '',
      // loginhistory: this.$store.state.LoginHistory,
      // loginarr: this.$store.state.Login,
      userId: '',
      userrole: '',
      Todaysale:0,
     
    }
  },
  methods: {
    // fetchId(){
    //   for (let i = 0; i < this.loginhistory.length; i++) {
    //     if (this.loginhistory[i].Log_out == 'In_Working') {
    //       this.userId = this.loginhistory[i].UserId;
    //       for (let j = 0; j < this.loginarr.length; j++) {
    //         if (this.userId == this.loginarr[j].UserId) {
    //           this.role = this.loginarr[j].Role;
    //           break;
    //         } 
    //       }
    //     } 
    //   } 
    // },   

  },
  beforeMount(){
    this.role=this.$store.state.UserDetails.Role;
    // console.log(this.role);
      // this.fetchId();
  },
  
}
</script>


<!-- TO FETCH USER ROLE -->

<!-- fetchId(){
  for(let i=0;i<this.loginhistory.length;i++){
    for(let j=0;j<this.loginarr.length;j++){
      if(this.loginhistory[i].Log_out=='null' && this.loginarr[j].UserId==this.loginhistory[i].UserId){
        this.role=this.loginarr[j].Role;
      }
    }
  }
} -->