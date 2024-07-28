<template>
    <div id="managerTS1">
        <v-container>

            <h3 class="blue--text darken-6--text font-weight-black text-center mt-5 display-1 text-uppercase">Today Sales</h3>
            <h1 class="grey--text darken-6--text font-weight-bold text-center mt-5 display-3">{{ MgTodaysales }}</h1>

        </v-container>


    </div>
</template>

<script>
export default {


    name: 'managerTS1',
    data() {
        return {
           

        }
    },
    props:{
        MgTodaysales:Number
    },
    // beforeMount() {
    //     // this.todaysales();

    // },
    methods: {
        todaydate() {
            this.currentdate = new Date().toJSON().slice(0, 10);
        },
        todaysales() {
            this.todaydate();
            for (let i = 0; i < this.BillMaster.length; i++) {
                if (this.currentdate == this.BillMaster[i].BillDate) {
                    this.todaytotal += this.BillMaster[i].NetPrice;

                }
            }
        },
    }
}


</script>