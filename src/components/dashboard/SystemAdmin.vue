<template>
 
      <v-container class="mt-16">
       
          <h1 class="blue--text darken-6--text font-weight-black text-center mt-16 display-2">Welcome !</h1>
          <h1 class="grey--text text--darken-1 font-weight-black text-center mt-16 display-2">{{ UserName.toUpperCase() }}</h1>
          <!-- {{ this.$store.state.Temp.user }} -->
        
          
      </v-container>

</template>

<script>

export default {
  name: "sysadmin",
  data(){
    return{
      UserName:"",
    }
  },
  beforeMount(){
    this.UserName=this.$store.state.UserDetails.User_id
  }
  
};
</script>
