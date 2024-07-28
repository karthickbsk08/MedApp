<template>
  <v-app >
    {{ this.$store.state.UserDetails.Role }}
       <!-- <Navmenubar v-show="this.$store.state.UserDetails.Role == '' ? false : true" :role="this.$store.state.Temp.role" /> -->
       <Navmenubar v-show="this.$store.state.UserDetails.Role == '' ? false : true" :role="this.$store.state.UserDetails.Role" />
     <!-- <v-main> -->
      <router-view />
     <!-- </v-main> -->
   </v-app>
</template>

<script>

import Navmenubar from '../src/components/login/Navmenubar.vue';
export default {
  name: "App",
  components: {
    Navmenubar,
    },
    props:{

    },
  data() {
    return {
      totaltime: '',
      userrole: '',
      navdata: false,
    }
  },methods: {
  
  }, 
}
</script>

