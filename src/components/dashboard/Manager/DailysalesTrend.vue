<template>
    <v-card elevation="2" max-width="650" class="pa-5 grey lighten-3">
        <!-- <h1 class="mb-2">DailysalesTrend</h1> -->
        <apexchart width="600" height="350" type="area" :options="chartOptions" :series="series"></apexchart>
        <!-- {{ MgDailySalesArr }} -->
    </v-card>
</template>

<script>
import VueApexCharts from 'vue-apexcharts';
export default {
    name: 'Dailytrend',
    components: {
        apexchart: VueApexCharts,
    },
    props: {
        MgDailySalesArr: []
    },
    data() {
        return {
            series: [{
                name: "Desktops",
                data: []
            }],
            chartOptions: {
                chart: {
                    height: 1000,
                    width: 500,
                    type: 'area',
                    zoom: {
                        enabled: false
                    }
                },
                dataLabels: {
                    enabled: true
                },
                stroke: {
                    curve: 'smooth'
                },
                title: {
                    text: 'Daily sales Trend',
                    align: 'center',
                    style: {
                        fontSize: "30px",
                        fontWeight: 700,
                        color: '#37474F',
                        textTransform: "uppercase",
                    }
                },
                grid: {
                    row: {
                        colors: ['#f3f3f3', 'transparent'], // takes an array which will be repeated on columns
                        opacity: 0.5
                    },
                },
                xaxis: {
                    categories: [],
                    title: {
                        text: 'Days'
                    }
                },
                yaxis: {
                    title: {
                        text: 'Amount'
                    },

                }
            },
            BillMaster: this.$store.state.BillMaster,
            todayDay: '',
            currentdate: '',
            TotalSales: [],
            date : [],
            amount : [],
            UserSales:0,
        }
    },
    methods: {


        getIndex() {
            

            for (let i = 0; i < 7; i++) {
                this.getDate(i);
            }

            this.series[0]= {
                data : this.date
            }

            this.chartOptions = {
                xaxis :{
                    categories : this.amount
                }
            }
            // console.log("DST");
            // console.log(this.date);
            // console.log(this.amount);
            

            
        },
        getDate(i) {

            this.UserSales=0

            let DateObj = new Date();
            DateObj.setDate(DateObj.getDate() - i);
            let TempDay = DateObj.getDay();
            let temp = true

            const days = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];

            for (let i = 0; i < this.MgDailySalesArr.length; i++) {

                if (days[TempDay] == this.MgDailySalesArr[i].day) {
                    this.UserSales = this.MgDailySalesArr[i].total_amount
                    temp = false
                }
            }

            if (temp == false) {

                this.date.unshift(this.UserSales);
                this.amount.unshift(days[TempDay])

            } else {
                this.date.unshift(this.UserSales);
                this.amount.unshift(days[TempDay])
            }

            
           

        },
        getTotalSales(from) {

            // var this.UserSales = 0

            for (let i = 0; i < this.BillMaster.length; i++) {
                if (this.BillMaster[i].BillDate == from) {

                    this.UserSales += this.BillMaster[i].NetPrice;
                }
            }
            // console.log(this.UserSales);
            return this.UserSales;

        },
        SetData(Amount, day) {
            const days = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];

            this.TotalSales.push({ Days: days[day], UserSales: Amount })

            // console.log(this.TotalSales);
        },
        MapData() {
            for (let i = 0; i < this.TotalSales.length; i++) {
                this.series[0].data.unshift(this.TotalSales[i].UserSales);
                this.chartOptions.xaxis.categories.unshift(this.TotalSales[i].Days);
            }
            // console.log(this.series[0].data);
            // console.log(this.chartOptions.xaxis.categories);
        },
        AssignData() {

        }
    },
    watch:{
        MgDailySalesArr(){
            
            this.getIndex();

        }
    },



    // mounted() {
       


    //     // this.MapData();

    //     // GETTING CURRENTDATE
    //     // this.currentdate = new Date().toJSON().slice(0, 10);

    //     // console.log(this.currentdate);

    //     //GETTING CURRENT DAY
    //     // const d = new Date().getDay();
    //     // console.log("No of Rotations" + d);

    //     //ROTATE ARRAY BASED ON THE DAY
    //     // function rotateArray(arr, k) {
    //     //     const length = arr.length;
    //     //     const effectiveRotation = k % length;

    //     //     const rotatedArr = [];

    //     //     for (let i = 0; i < length; i++) {
    //     //         rotatedArr[(i + effectiveRotation) % length] = arr[i];
    //     //     }

    //     //     return rotatedArr;
    //     // }

    //     // Example usage
    //     // const days = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
    //     // const rotatedDays = rotateArray(days, 6 - d);


    //     // this.chartOptions.xaxis.categories = rotatedDays;
    //     // console.log(rotatedDays);
    //     // console.log(this.chartOptions.xaxis.categories);

    //     //GETTING THE ALL DATES FROM BILL MASTER

    //     // let arr = [];
    //     // for (let i = 0; i < this.BillMaster.length; i++) {
    //     //     if (this.BillMaster[i].BillDate <= this.currentdate) {

    //     //         // console.log(this.BillMaster[i].BillDate);
    //     //         arr.push(this.BillMaster[i].BillDate);

    //     //     }

    //     // }

    //     // console.log(arr);
    //     //SORT IT BASED ON DATE
    //     // arr = arr.sort();
    //     // console.log(arr);
    //     // let datecount=0;
    //     // let newdatearr=[];

    //     // for (let w = 0; w < arr.length-1; w++) {
    //     //     if(arr[w]!=arr[w+1]){
    //     //        newdatearr.push(arr[w]);
    //     //     }
    //     // }
    //     // newdatearr.push(arr[arr.length-1]);
    //     // console.log(newdatearr);



    //     //GETTING LAST SEVEN DATES FROM LAST TO 7 DAYS BEFORE
    //     // let array1 = [];
    //     // for (let j = newdatearr.length - 1; j >= newdatearr.length - 7; j--) {
    //     //     array1.push(newdatearr[j])
    //     // }
    //     // console.log(array1);

    //     //CREATE A FINAL ARRAY AND NETPRICE IS STORED THERE
    //     // let NewArray = [];
    //     // let temparr = [];
    //     // let TotalPrice = 0;

    //     // for (let k = 0; k < array1.length; k++) {
    //     //     for (let l = 0; l < this.BillMaster.length; l++) {
    //     //         if (array1[k] == this.BillMaster[l].BillDate) {
    //     //             temparr.push(this.BillMaster[l].NetPrice);

    //     //         }


    //     //     }
    //     //     for (let n = 0; n < temparr.length; n++) {
    //     //         TotalPrice += temparr[n];
    //     //     }
    //     //     NewArray.unshift(TotalPrice);
    //     //     temparr = [];
    //     //     TotalPrice = 0;
    //     // }
    //     // this.series[0].data = NewArray;






    // },
}

</script>

<!-- // data: {

    //   series: [{
    //       name: "Desktops",
    //       data: [10, 41, 35, 51, 49, 62, 69, 91, 148]
    //   }],
    //   chartOptions: {
    //     chart: {
    //       height: 350,
    //       type: 'line',
    //       zoom: {
    //         enabled: false
    //       }
    //     },
    //     dataLabels: {
    //       enabled: false
    //     },
    //     stroke: {
    //       curve: 'straight'
    //     },
    //     title: {
    //       text: 'Product Trends by Month',
    //       align: 'left'
    //     },
    //     grid: {
    //       row: {
    //         colors: ['#f3f3f3', 'transparent'], // takes an array which will be repeated on columns
    //         opacity: 0.5
    //       },
    //     },
    //     xaxis: {
    //       categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep'],
    //     }
    //   }, -->


    <!-- // const weekday = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];

    // let day = weekday[d.getDay()];
    // this.options.xaxis.categories.unshift(day);
    // console.log(weekday);

    // for(let i=0;i<4;i++){
    //     let first=weekday[weekday.length-1];
    //     for(let j=0;j<weekday.length-1;j++){
    //         weekday[j+1]=weekday[j];

    //     }
    //     weekday[0]=first;
    // }
    // console.log(weekday);

    // console.log(day);

    // this.todayDay = new Date().getDay;
    // console.log(this.todayDay);
    // this.billdate=this.BillMaster[this.BillMaster.length-1].BillDate;
    // this.billdate.
    // for(let i=rotatedDays.length-1;i>=0;i--){
    //     this.options.xaxis.categories.push(rotatedDays[i]);
    // }  getDate(i) {

            let DateObj = new Date();
            DateObj.setDate(DateObj.getDate() - i);

            var Month = DateObj.getMonth() + 1;

            let TempDay = DateObj.getDay();
            // console.log(TempDay);

            if (Month <= 9) {
                Month = '0' + Month
            }

            var Year = DateObj.getFullYear();

            var date = DateObj.toString().slice(8, 10);

            var from = `${Year}-${Month}-${date}`;
            // console.log(from);

            // this.getTotalSales(from)

            this.SetData(this.getTotalSales(from), TempDay)


        },-->