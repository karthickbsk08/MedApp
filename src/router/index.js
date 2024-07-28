import Vue from "vue";
import VueRouter from "vue-router";
import login from "../views/Login.vue";
import dashboard from '../views/Dashboard.vue';
import loginhistory from "../views/LoginHistory.vue";
import adduser from "../views/AddUser.vue";
import stockview from '../views/StockView.vue';
import stockentry from '../views/StockEntry.vue';
import billentry from '../views/BillEntry.vue';
import salesreport from '../views/SalesReport.vue'; 





Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Login",
    component: login,
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: dashboard,
  },
  {
    path: "/loginhistory",
    name: "Loginhistory",
    component: loginhistory,
  },
  {
    path: "/adduser",
    name: "Adduser",
    component: adduser,
  },
  {
    path: "/stockview",
    name: "Stockview",
    component: stockview,
  },
  {
    path: "/stockentry",
    name: "Stockentry",
    component: stockentry,
  },
  {
    path: "/billentry",
    name: "Billentry",
    component: billentry,
  },
  {
    path: "/salesreport",
    name: "Salesreport",
    component: salesreport,
  },  
];

const router = new VueRouter({
  routes,
});

export default router;




























// {
//   path: "/loginhistory",
//   name: "loginhistory",
//   // route level code-splitting
//   // this generates a separate chunk (about.[hash].js) for this route
//   // which is lazy-loaded when the route is visited.
//   component: () =>
//     import(/* webpackChunkName: "about" */ "../views/About.vue"),
// },
