<template>
    <v-card max-width="650" class="pa-5 grey lighten-3">
        <!--  -->
        <!-- <h1 >Monthly Sales Trend</h1> -->
        <apexchart type="line" height="350" width="600" :options="chartOptions" :series="series"></apexchart>
        <!-- {{ MgMonthlysalesArr }} -->

    </v-card>
</template>

<script>

import VueApexCharts from "vue-apexcharts";
export default {
    name: 'Monthlysales',
    components: {
        apexchart: VueApexCharts,
    },
    props: {
        MgMonthlysalesArr: []
    },
    data() {
        return {

            series: [
                {
                    name: "sales",
                    data: []
                },

            ],
            chartOptions: {
                chart: {
                    height: 350,
                    type: 'line',
                    dropShadow: {
                        enabled: true,
                        color: '#000',
                        top: 18,
                        left: 7,
                        blur: 10,
                        opacity: 0.2
                    },
                    toolbar: {
                        show: true
                    }
                },
                colors: ['#77B6EA', '#545454'],
                dataLabels: {
                    enabled: true,
                },
                stroke: {
                    curve: 'smooth'
                },
                title: {
                    text: 'Monthly sales trend',
                    align: 'center',
                    fontSize: "25px",
                },
                grid: {
                    borderColor: '#e7e7e7',
                    row: {
                        colors: ['#f3f3d4', 'transparent'], // takes an array which will be repeated on columns
                        opacity: 0.5
                    },
                },
                markers: {
                    size: 1
                },
                xaxis: {
                    categories: [],
                    title: {
                        text: 'Month',
                        size: 10,
                    }
                },
                yaxis: {
                    title: {
                        text: 'Amount',
                        fontSize: "25px",
                    },
                    legend: {
                        position: 'top',
                        horizontalAlign: 'right',
                        floating: false,
                        offsetY: -25,
                        offsetX: -5
                    }
                },



            },
            Month: [],
            MonthlySales: [],
            UserSales: 0,
        }
    },
    watch: {
        MgMonthlysalesArr() {
            this.SetData1()
        }
    },

    methods: {
        GetIndex() {


            for (let i = 0; i < 12; i++) {
                this.GetMonth(i);
            }
        },
        GetMonth(index) {

            let DateObj = new Date()


            DateObj.setMonth(DateObj.getMonth() - index);
            console.log("Date", DateObj);


            // var Month = DateObj.getMonth() +1;

            var TempMonth = DateObj.getMonth();
            var mL = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
            console.log("data", mL[TempMonth]);
            // console.log("Month" + );

            // var Year = DateObj.getFullYear();
            // // console.log("Year  : " + Year);



        },

        getTotalSales(from, To) {
            // console.log(from);
            // console.log(To);
            var UserSales = 0

            for (let i = 0; i < this.BillMaster.length; i++) {
                if (this.BillMaster[i].BillDate >= from && this.BillMaster[i].BillDate <= To) {
                    // console.log(this.BillMaster[i].BillDate);
                    UserSales += this.BillMaster[i].NetPrice;
                }
            }
            // console.log("Usersales "+UserSales);
            return UserSales;


        },
        SetData(Amount, month) {
            // var mL = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
            var mL = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
            this.MonthlySales.push({ Month: mL[month], UserSales: Amount });
            // console.log(this.MonthlySales);

        },
        PushArray() {

            for (var j = 0; j < this.MonthlySales.length; j++) {
                this.series[0].data.unshift(this.MonthlySales[j].UserSales);
                this.chartOptions.xaxis.categories.unshift(this.MonthlySales[j].Month);
            }
            // console.log(this.series[0].data);
            // console.log(this.chartOptions.xaxis.categories);
        },
        SetData1() {
            const monthNames = [
                "Jan", "Feb", "Mar", "Apr", "May", "Jun",
                "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
            ];
            // Get the current date
            const currentDate = new Date();

            // Loop through the last 12 months
            for (let i = 0; i < 12; i++) {
                var temp = true
                this.UserSales = 0
                const monthIndex = (currentDate.getMonth() - i + 12) % 12;
                const monthName = monthNames[monthIndex];

                for (let i = 0; i < this.MgMonthlysalesArr.length; i++) {

                    if (monthName == this.MgMonthlysalesArr[i].month) {
                        this.UserSales = this.MgMonthlysalesArr[i].total_amount;
                        // console.log(this.MgMonthlysalesArr[i].month,UserSales);
                        temp = false
                    }
                }
                if (temp == false) {
                    this.Month.unshift(monthName)
                    this.MonthlySales.unshift(this.UserSales)

                } else {
                    this.Month.unshift(monthName)
                    this.MonthlySales.unshift(this.UserSales)
                }
            }
            this.chartOptions = {
                xaxis: {
                    categories: this.Month
                }
            }
            this.series[0] = {
                data: this.MonthlySales
            }
            // console.log("Month");
            // console.log(this.Month);
            // console.log(this.MonthlySales);
        }
    }
}
</script>