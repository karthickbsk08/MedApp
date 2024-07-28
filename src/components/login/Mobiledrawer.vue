<template>
  <div class="d-inline-block" >
    <v-menu rounded="lg" offset-y>
      <template v-slot:activator="{ attrs, on }">
        <!-- <v-btn text class="grey lighten-2 font-weight-black black--text " v-bind="attrs" v-on="on"> DROP DOWN -->

        <v-app-bar-nav-icon  class="grey darken-1" v-bind="attrs" v-on="on"></v-app-bar-nav-icon>
        <!-- </v-btn> -->
      </template>

      <v-list dense >
        <v-list-item v-for="item in array" :key="item.title" :to="item.route"   >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>

<!-- <script>

export default {
    data() {
        return {
            items: [
                { title: 'LOG OUT', icon: 'mdi-logout-variant', route: '/' },
                { title: 'LOGIN HISTORY', icon: 'mdi-login', route: '/loginhistory' },
                { title: 'ADD USER', icon: 'mdi-account', route: '/adduser' },
                { title: 'STOCK VIEW', icon: 'mdi-teamviewer', route: '/stockview' },
                { title: 'STOCK ENTRY', icon: 'mdi-basket-fill', route: '/stockentry' },
                { title: 'BILL ENTRY', icon: 'mdi-receipt', route: '/billentry' },
                { title: 'SALES REPORT', icon: 'mdi-sale', route: '/salesreport' },
            ],
        }
    }
}
</script> -->


<script>
export default {
  name: "",
  props: {
    Role: String,
  },
  data: () => ({
    SysAdmin: [
      { title: 'DASHBOARD', icon: 'mdi-view-dashboard', route: '/dashboard' },
      { title: 'ADD USER', icon: 'mdi-account', route: '/adduser' },
      { title: 'LOGIN HISTORY', icon: 'mdi-login', route: '/loginhistory' },
    ],
    BLR: [
      { title: 'DASHBOARD', icon: 'mdi-view-dashboard', route: '/dashboard' },
      { title: 'BILL ENTRY', icon: 'mdi-receipt', route: '/billentry' },
      { title: 'STOCK VIEW', icon: 'mdi-teamviewer', route: '/stockview' },
    ],
    INV: [
      { title: 'DASHBOARD', icon: 'mdi-view-dashboard', route: '/dashboard' },
      { title: 'STOCK VIEW', icon: 'mdi-teamviewer', route: '/stockview' },
      { title: 'STOCK ENTRY', icon: 'mdi-basket-fill', route: '/stockentry' },
    ],
    MGR: [
      { title: 'DASHBOARD', icon: 'mdi-view-dashboard', route: '/dashboard' },
      { title: 'STOCK ENTRY', icon: 'mdi-basket-fill', route: '/stockentry' },
      { title: 'STOCK VIEW', icon: 'mdi-teamviewer', route: '/stockview' },
      { title: 'SALES REPORT', icon: 'mdi-sale', route: '/salesreport' },
    ],
    Roledata: "System Admin",
    array: [],
  }),
  beforeMount() {
    this.Roledata = this.Role;
  },
  watch: {
    Role() {
      if (this.Role == 'System Admin') {
        this.array = this.SysAdmin;
       
      } else if (this.Role == "Inventry") {
        this.array = this.INV;

      } else if (this.Role == "Biller") {       
        this.array = this.BLR;      

      } else if (this.Role == "Manager") {
        this.array = this.MGR;
      }
    }
  }
}
</script>