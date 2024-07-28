<template>
 
       <v-card max-width="650" class="pa-5 grey lighten-3" >
        <!-- <h1 class="mb-2">Today Salesman Performance</h1> -->
        <apexchart type="donut" width="600" height="350" :options="chartOptions" :series="series"></apexchart>
        <!-- {{ MgTdySalesPefArr }} -->
  
       </v-card>
</template>


<script>
import VueApexCharts from 'vue-apexcharts';
export default {
    name: "Todaysalesmanperformance",
    components: {
        apexchart: VueApexCharts,
    },
    props:{
        MgTdySalesPefArr:[],
    },
    data() {
        return {
            series: [],
            chartOptions: {
                chart: {
                    width: 380,
                    type: 'donut',
                },
                plotOptions: {
                    pie: {
                        startAngle: -90,
                        endAngle: 270
                    }
                },
                dataLabels: {
                    enabled: true
                },
                fill: {
                    type: 'gradient',
                },
                // colors: ['#F44336', '#E91E63', '#9C27B0'],
                labels: [],
                legend: {
                    formatter: function (val, opts) {
                        return val + " - " + opts.w.globals.series[opts.seriesIndex]
                    }
                },
               
                title: {
                    text: 'Today Salesman Performance',
                    align:'center',
                },
                responsive: [{
                    breakpoint: 480,
                    options: {
                        chart: {
                            width: 200
                        },
                        legend: {
                            position: 'bottom'
                        }
                    }
                }]
            },
            BillMaster: this.$store.state.BillMaster,
            currentdate: '',
            Todaysales: [],
            roleArr: [],
            FinalSales: [],
        }
    },
    watch:{
        MgTdySalesPefArr(){
            this.SetData()
        }
    },
    methods: {
        getDate() {
            this.currentdate = new Date().toJSON().slice(0, 10);
        },
        getBills() {
            for (let i = 0; i < this.BillMaster.length; i++) {
                if (this.currentdate == this.BillMaster[i].BillDate) {
                    // console.log(this.BillMaster[i]);
                    this.Todaysales.push(this.BillMaster[i]);
                }
            }
            this.Todaysales.sort();
            // console.log(this.Todaysales);
        },
        getUserSales() {
            let Usersales = 0;
            let CurrentUser = "";

            for (let k = 0; k < this.roleArr.length; k++) {
                CurrentUser = this.roleArr[k].split("@")[0];
                for (let j = 0; j < this.Todaysales.length; j++) {

                    if (this.roleArr[k] == this.Todaysales[j].UserId) {
                        Usersales += this.Todaysales[j].NetPrice;
                    }
                }
                this.FinalSales.push({ UserId: CurrentUser, Usersales: Usersales });
                CurrentUser = "";
                Usersales = 0;


            }

            // console.log(this.FinalSales);

        },
        getRole() {
            let Role = this.$store.state.Login;
            for (let k = 0; k < Role.length; k++) {
                if (Role[k].Role == "Biller") {
                    this.roleArr.push(Role[k].UserId)
                }

            }
            // console.log(this.roleArr);
        },
        SetData() {
            for (let i = 0; i < this.MgTdySalesPefArr.length; i++) {
                this.series.push(this.MgTdySalesPefArr[i].total_amount);
                this.chartOptions.labels.push(this.MgTdySalesPefArr[i].user_id);
            }
        }

    }
}
</script>