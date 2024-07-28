
<template>
 
    <!-- <v-btn
     
      @click="snackbar = true"
    >
      Open Snackbar
    </v-btn> -->
    <v-snackbar v-model="obj.snackbar" :top="obj.topvar" :bottom="obj.bottomvar" :color="obj.color" rounded
      class="font-weight-black text-center mt-16 black--text" :timeout="3000" shaped :location="obj.location">
      <span class="white--text font-weight-black text-center"> {{ obj.text }}
      </span>
      <template v-slot:action="{ attrs }">
        <v-btn dark class="white--text font-weight-black" text v-bind="attrs" @click="obj.snackbar = false" rounded>
          Close
        </v-btn>
      </template>
    </v-snackbar>
  
</template>


<script>
export default {
  name: 'snackbar',
  props: {
    obj: {}
  }
}
</script>

