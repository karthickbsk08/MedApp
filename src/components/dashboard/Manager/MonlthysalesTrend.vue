<template>
    <v-card class="pa-5 grey lighten-3" max-width="650">
        <!-- <h1 >CurrentMonlthsalesmanperformance</h1> -->
        <apexchart type="bar" height="350" width="600" :options="chartOptions" :series="series"></apexchart>

        <!-- {{ MgCurrentMonArr }} -->
    </v-card>
</template>
<script>
import VueApexCharts from 'vue-apexcharts';
export default {
    name: 'CurrentMonlthsalesmanperformance',
    components: {
        apexchart: VueApexCharts,
    },
    props: {
        MgCurrentMonArr: []
    },
    data() {
        return {
            series: [{
                data: []
            },
           ],
            chartOptions: {
                chart: {
                    height: 350,
                    type: 'bar',
                    events: {
                        // click: function(chart, w, e) {
                        //   console.log(chart, w, e)
                        // }
                    }
                },
                title: {
                    text: 'Current Monlth salesman performance',
                    align: 'center',
                    fontSize: '20'
                },
                // colors: colors,
                plotOptions: {
                    bar: {
                        columnWidth: '45%',
                        distributed: true,
                    }
                },
                dataLabels: {
                    enabled: true,
                },
                legend: {
                    show: false,
                },
                xaxis: {
                    categories: [


                    ],
                    labels: {
                        style: {
                            // colors: colors,
                            fontSize: '12px'
                        }
                    },
                    title: {
                        text: 'SalesMan',
                        fontSize: '20px'
                    }
                },
                yaxis: {
                    title: {
                        text: 'Amount'
                    },

                }
            },
            // BillMaster: this.$store.state.BillMaster,
            // currentmonth: '',
            // CurrentSales: [],
            // roleArr: [],
            TotalSales: [],
            User: [],

        }
    },
    watch: {
        MgCurrentMonArr() {
            this.setData()
        }
    },
    methods: {
        getDate() {
            this.currentmonth = new Date().toJSON().slice(0, 7);
            // console.log(this.currentmonth);
        },
        getSalesMonth() {
            for (let i = 0; i < this.BillMaster.length; i++) {
                if (this.BillMaster[i].BillDate.slice(0, 7) == this.currentmonth) {
                    this.CurrentSales.push(this.BillMaster[i]);

                }

            }
            // console.log(this.CurrentSales);
            // this.CurrentSales.sort();
            // console.log(this.CurrentSales);
        },
        getUserSales() {
            let UserSales = 0;
            let currentuser = "";
            // for (let j = 0; j < this.CurrentSales.length; j++) {
            //     UserSales=0;
            //   for (let k = 0; k < this.roleArr.length; k++) {
            //         if(this.CurrentSales[j].UserId==this.roleArr[k]){
            //             UserSales+=this.CurrentSales[j].NetPrice;


            //         }

            //   }


            // }
            // this.TotalSales.push({UserId:this.CurrentSales[j].UserId,UserSales:UserSales});
            // console.log(this.TotalSales);




            for (let j = 0; j < this.roleArr.length; j++) {
                currentuser = this.roleArr[j].split("@")[0];
                for (let k = 0; k < this.CurrentSales.length; k++) {
                    if (this.roleArr[j] == this.CurrentSales[k].UserId) {
                        UserSales += this.CurrentSales[k].NetPrice;
                    }
                }
                this.TotalSales.push({ UserId: currentuser, UserSales: UserSales });
                UserSales = 0;
                currentuser = "";
            }
            this.TotalSales.sort();
            // console.log(this.TotalSales);

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
        setData() {
            // console.log("CMPS");
            for (let n = 0; n < this.MgCurrentMonArr.length; n++) {
                this.User.push(this.MgCurrentMonArr[n].user_id);
                this.TotalSales.push(this.MgCurrentMonArr[n].total_amount);
            }
            // console.log("CMPE");

            this.chartOptions = {
                xaxis: {
                    categories: this.User
                }
            }
            this.series[0] = {
                data: this.TotalSales
            }
            
            // console.log("CMPRESP");
            // console.log(this.chartOptions.xaxis.categories);
            // console.log(this.series[0].data);

        }    }


}
</script>