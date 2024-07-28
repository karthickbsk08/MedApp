<template>
    <v-container :fluid="true">
        <!-- {{ TodaySales1 }}<br><br>
        {{ Inventoryvalue1 }}<br><br>
        {{ DailysalesTrend1 }}<br><br>
        {{ MonthlySales1 }}<br><br>
        {{ Todaysalesperformance1 }}<br><br>
        {{ MonlthysalesTrend1 }}<br><br> -->
        <v-layout row wrap class="justify-space-around ml-15 mt-5 ">
            <v-flex md-6 class="mb-10">
                <TodaySales :MgTodaysales="TodaySales1" />

            </v-flex>
            <v-flex md-6 class="mb-10">
                <Inventoryvalue :MgInventoryValue="Inventoryvalue1" />
            </v-flex>

        </v-layout>
        <v-layout row wrap class="justify-space-around ml-15 mt-5 ">
            <v-flex md-5 sm-12 class="mb-10">
                <DailysalesTrend :MgDailySalesArr="DailysalesTrend1" />
            </v-flex>
            <v-flex md-5 sm-12 class="mb-10">
                <MonthlySales :MgMonthlysalesArr="MonthlySales1" />
            </v-flex>
        </v-layout>

        <v-layout row wrap class="justify-space-around ml-15 mt-5 ">
            <v-flex md-5 sm-6 class="mb-10">
                <Todaysalesperformance :MgTdySalesPefArr="Todaysalesperformance1" />
            </v-flex>
            <v-flex md-5 sm-6>
                <MonlthysalesTrend :MgCurrentMonArr="MonlthysalesTrend1" />
            </v-flex>
        </v-layout>
       
    </v-container>
</template>


<script>
import EventService from '../../../Services/EventService';
import TodaySales from './MgTodaysales.vue';
import Inventoryvalue from './inventryvalue.vue';
import DailysalesTrend from '../Manager/DailysalesTrend.vue';
import MonthlySales from '../Manager/MonthlySales.vue';
import Todaysalesperformance from './TodaySalesmanperformance.vue';
import MonlthysalesTrend from './MonlthysalesTrend.vue';
export default {
    name: 'managerdash',
    data() {
        return {
            FinalResult: {},
            TodaySales1: 0,
            Inventoryvalue1: 0,
            DailysalesTrend1: [],
            MonthlySales1: [],
            Todaysalesperformance1: [],
            MonlthysalesTrend1: [],
        }
    },
    components: {
        TodaySales,
        Inventoryvalue,
        DailysalesTrend,
        MonthlySales,
        Todaysalesperformance,
        MonlthysalesTrend,
    },
    beforeMount() {
        EventService.ManagerDash()
            .then((response) => {
                console.log(response);
                // console.log(response.data);
                // this.FinalResult = response.data
                if (response.data.status == "S") {

                    // if (response.data.oatoday_Sales != 0){
                    //     this.TodaySales1 = response.data.oatoday_Sales

                    // }else if (response.data.inventory_Value != 0){
                    //     this.Inventoryvalue1 = response.data.inventory_Value
                    // }else if (response.data.dailysalesarr != null){
                    //     this.DailysalesTrend1 = response.data.dailysalesarr
                    // }else if (response.data.monthlysalesarr !=null){
                    //     this.MonthlySales1 = response.data.monthlysalesarr
                    // }else if (response.data.todaysalesarr != null){
                    //     this.Todaysalesperformance1 = response.data.todaysalesarr
                    // }else if (response.data.currentmontharr !=null){
                    //     this.MonlthysalesTrend1 = response.data.currentmontharr
                    // }

                    this.TodaySales1 = response.data.oatoday_Sales
                    this.Inventoryvalue1 = response.data.inventory_Value
                    this.DailysalesTrend1 = response.data.dailysalesarr
                    this.MonthlySales1 = response.data.monthlysalesarr
                    this.Todaysalesperformance1 = response.data.todaysalesarr
                    this.MonlthysalesTrend1 = response.data.currentmontharr

                } else if (response.data.status == "E") {
                    console.log(response.data.msg);
                }

            })
            .catch((error) => {
                console.log(error);
            })
    }

}

</script>