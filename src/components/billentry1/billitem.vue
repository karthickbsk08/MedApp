<template>
    <v-container class="pb-2 ">
        <snackbar :obj="obj" />
        <v-expansion-panels>
            <v-expansion-panel class="grey lighten-3">
                <v-expansion-panel-header class="blue lighten-2 white--text font-weight-bold" elevation="0" dense>
                    Item
                </v-expansion-panel-header>
                <v-expansion-panel-content class=" pl-6 " elevation="0">
                    <v-layout class="mt-10">
                        <v-flex>
                            <v-autocomplete :items="item" label="Medicine Name" outlined dense
                                v-model="MedicineDetails.medicine_name" @click="StockReceive"></v-autocomplete>

                        </v-flex>
                        <v-flex>
                            <v-text-field v-model.number="MedicineDetails.quantity" label="Quantity" outlined class="pl-2"
                                type="number" dense></v-text-field>
                        </v-flex>
                        <v-flex>
                            <v-btn dark class="blue font-weight-bold ml-6" dense @click="Addbtn">
                                ADD
                            </v-btn>
                        </v-flex>
                    </v-layout>
                </v-expansion-panel-content>
            </v-expansion-panel>
        </v-expansion-panels>
        <!-- {{ LocalSArr }}<br><br>
            {{ stockarr }}  -->
        <!-- <br><br>{{ RespStock }} <br><br> -->
        <!-- {{ AddArr }} -->
    </v-container>
</template>
  
<script>

import EventService from '../../Services/EventService.js'
import snackbar from '../snackbar.vue';

export default {
    name: "billitem",
    components: {
        snackbar,
    },
    props: {
        roughdata: String,
        RespStock: [],
        itemArr: [],
    },
    data() {
        return {
            MedicineDetails: {
                medicine_name: "",
                quantity: 0,
            },

            item: [],
            unitprice: 0,
            AddArr: [],
            tempmedname: '',
            tempunitprice: 0,
            amount: 0,
            totalprice: 0,
           
          
            obj: {
                snackbar: false,
                color: 'green',
                topvar: false,
                bottomvar: false,
                text: '',
            },
            Resp: [],



        }
    },
    watch: {
        itemArr() {
            this.AddArr = []
        }
    },
    methods: {
        StockReceive() {           
            for (let i = 0; i < this.RespStock.length; i++) {
                this.item.push(this.RespStock[i].medicine_name);
            }          
        },
        // addbtn() {
        //     this.qty = Math.round(this.qty);
        //     // console.log("Add start");
        //     // console.log(this.LocalSArr,this.LocalSArr[0].quantity);
        //     // console.log(this.stockarr);

        //     let temp = false;
        //     if (this.medicinename != '' && this.qty != '' && this.qty > 0) {
        //         for (let i = 0; i < this.LocalSArr.length; i++) {
        //             if (this.LocalSArr[i].MedicineName == this.medicinename && this.qty != 0 && this.qty > 0) {
        //                 // let TempQuantity = this.LocalSArr[i].quantity;

        //                 if (this.qty <= this.LocalSArr[i].quantity) {
        //                     this.tempmedname = this.medicinename;
        //                     this.tempunitprice = this.LocalSArr[i].Unitprice;
        //                     // TempQuantity=TempQuantity-this.qty;
        //                     this.LocalSArr[i].quantity -= this.qty;
        //                     temp = true;
        //                 }
        //                 // temp=false;
        //             }

        //         }

        //         if (temp == true) {
        //             this.brandname();
        //             this.EachAmount();
        //             this.AddArr.push({ medicinename: this.medicinename, Brand: this.brand, Qty: this.qty, Amount: this.amount, Unitprice: this.tempunitprice })
        //             this.totalprice += this.amount;


        //             this.medicinename = '';
        //             this.qty = '';


        //         } else {
        //             // alert('There is no enough quantity available in inventry');
        //             this.obj = {
        //                 snackbar: true,
        //                 color: 'red',
        //                 topvar: true,
        //                 bottomvar: false,
        //                 text: 'There is no enough quantity available in inventry',
        //             };
        //             temp = true;
        //             this.medicinename = '';
        //             this.qty = '';
        //         }

        //     } else {
        //         // alert('Enter the proper Data');
        //         this.obj = {
        //             snackbar: true,
        //             color: 'red',
        //             topvar: true,
        //             bottomvar: false,
        //             text: 'Enter the valid inputs',
        //         };


        //     }
        // },
        Addbtn() {
            var temp = false
            if (this.MedicineDetails.medicine_name != '' && this.MedicineDetails.quantity != '' && this.MedicineDetails.quantity > 0) {
                this.MedicineDetails.quantity = Math.round(this.MedicineDetails.quantity)

                for (let i = 0; i < this.RespStock.length; i++) {
                    if (this.RespStock[i].medicine_name == this.MedicineDetails.medicine_name && this.RespStock[i].quantity >= this.MedicineDetails.quantity) {
                        this.RespStock[i].quantity -= this.MedicineDetails.quantity
                        temp = true;
                    }
                }
                if (temp == true) {
                    EventService.AddBill(this.MedicineDetails)
                        .then((response) => {
                            console.log(response);
                            if (response.data.status == "S") {
                                this.AddArr.push(response.data.medicinearr[0])
                                this.totalprice = this.totalprice + response.data.medicinearr[0].amount
                                this.$emit('EachNewMed', this.AddArr, this.totalprice);
                                this.MedicineDetails.medicine_name = ""
                                this.MedicineDetails.quantity = ""

                            } else if (response.data.status == "E") {
                                console.log(response.data.msg);
                                console.log(response.data.errmsg);
                            }
                        })
                        .catch((error) => {
                            console.log(error);
                        })

                    temp = false;
                } else {
                    this.obj = {
                        snackbar: true,
                        color: 'red',
                        topvar: false,
                        bottomvar: true,
                        text: 'There is no enough quantity available in inventry'
                    };
                    this.MedicineDetails.medicine_name = ""
                    this.MedicineDetails.quantity = ""
                }
            } else {
                // alert('Enter the proper Data');
                this.obj = {
                    snackbar: true,
                    color: 'red',
                    topvar: true,
                    bottomvar: false,
                    text: 'Enter the valid inputs',
                };
                this.MedicineDetails.medicine_name = ""
                this.MedicineDetails.quantity = ""
            }
        },




        // let temp = false;
        // if (this.medicinename != '' && this.qty != '' && this.qty > 0) {
        //     for (let i = 0; i < this.stockarr.length; i++) {
        //         if (this.stockarr[i].MedicineName == this.medicinename && this.qty != 0 && this.qty > 0) {
        //             let TempQuantity = this.stockarr[i].quantity;

        //             if (this.qty <= TempQuantity) {
        //                 this.tempmedname = this.medicinename;
        //                 this.tempunitprice = this.stockarr[i].Unitprice;
        //                 TempQuantity=TempQuantity-this.qty;
        //                 // this.stockarr[i].quantity-=this.qty;
        //                 temp = true;
        //             }
        //             // temp=false;



        //         }

        //     }

        //     if (temp == true) {
        //         this.brandname();
        //         this.EachAmount();
        //         this.AddArr.push({ medicinename: this.medicinename, Brand: this.brand, Qty: this.qty, Amount: this.amount, Unitprice: this.tempunitprice })
        //         this.totalprice += this.amount;
        //         this.$emit('eachmedamount', this.AddArr, this.totalprice);

        //         this.medicinename = '';
        //         this.qty = '';


        //     } else {
        //         // alert('There is no enough quantity available in inventry');
        //         this.obj = {
        //             snackbar: true,
        //             color: 'red',

        //             topvar: false,
        //             bottomvar: true,
        //             text: 'There is no enough quantity available in inventry',
        //         };
        //         temp = true;
        //         this.medicinename = '';
        //         this.qty = '';
        //     }

        // } else {
        //     // alert('Enter the proper Data');
        //     this.obj = {
        //         snackbar: true,
        //         color: 'red',

        //         topvar: false,
        //         bottomvar: true,
        //         text: 'Enter the valid inputs',
        //     };
        // }
        // console.log("Add End");
        // console.log(this.LocalSArr);
        // console.log(this.stockarr);

        //     brandname() {
        //         // console.log('a');
        //         for (let i = 0; i < this.arr.length; i++) {
        //             if (this.medicinename == this.arr[i].MedicineName) {
        //                 this.brand = this.arr[i].Brand;
        //             }
        //         }
        //         // console.log(this.brand);
        //     },
        //     EachAmount() {
        //         this.amount = this.qty * this.tempunitprice;
        //     },
        // getLocalArray() {

        //     for (let i = 0; i < this.$store.state.BillStockDetails.length; i++) {
        //         this.LocalSArr.push({
        //             medicine_name: this.$store.state.BillStockDetails.MedicineName,
        //             quantity: this.$store.state.BillStockDetails.quantity,
        //         });
        //     }

        // }

        // },
    },
    


}
</script>
  