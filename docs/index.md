 <!DOCTYPE html>
<meta charset="UTF-8">
<style>

.node circle {
  fill: #fff;
  stroke: steelblue;
  stroke-width: 3px;
}

.node text {
  font: 12px sans-serif;
}

.link {
  fill: none;
  stroke: #ccc;
  stroke-width: 2px;
}

</style>

<body>

<!-- load the d3.js library -->
<script src="https://d3js.org/d3.v4.min.js"></script>
<script>

var treeData =
  {"name":"Rangu","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Akkamma","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Babi","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Chandravathi","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Dhavalesh","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Girishdhar","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Heera","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Yatishdhar","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Yogishdhar","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Devayani","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Mamta","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Vivek","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Dilipraj","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Mridulla","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Ujwala","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Harini","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Harshalatha","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Nithin","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Padhmalatha","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Lokanath","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Prajna","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Madhavi","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Bipin","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Dinnu","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sachin","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Prasanna","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Suneet","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Vineet","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Venkatramana","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Dhanaraj","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Gautham","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Kollu","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Devu","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Suraj","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Gopi","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Sabitha","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Venugopal","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Mohini","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Girish","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Tejpal","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Hithesh","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Yeshoda","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Mahabala","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Bhavani","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Jayaraj","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Padmaraja","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Roopalatha","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Tilothama","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sanki","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Daba","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"B","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Mamatha","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Santhosh","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Gururaj","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Deepali","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Deeya","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Harshali","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Madhava","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Vani","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Veena","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sharada","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Gautham","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Leka","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sheela","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sundari","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Chandraprabha","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"B1","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"B2","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Daskshayani","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"G","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Rahul","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Geetha","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Chitra","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Prashant","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Veena","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Godavari","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Chandini","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sandhya","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Premchand","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"B","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"G","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Preveena","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"B1","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"B2","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Raychand","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Pradeep","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Prajwal","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Pratap","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Vagdevi","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Deepak","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Divya","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Chaudappa","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Achuthananda","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Chandra Shekara","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Dayavathi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Devaki","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Jahnavi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Purandara","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Ramesha","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Shiva Shankara","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Koragappa","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Chinnapa","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Gangadhar","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Indira","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Kamalaksha","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sadashiva","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Vimala","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Yashoda","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Shesha sahukar","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Kamala","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Dhanalaxmi Purandara","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Heera","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Muthaakka","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Bhujanga banuvathi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Devdas purushotham","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Nagaveni","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Damodhara","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Dinakar","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"G","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Keshava","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Kunnudimi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Lalaji","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Tharachuda","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Neelamma","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sharada","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"vidya","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Tungamma","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Ananda","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Dinakar","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Divakar","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"G","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Gayathri","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Geetha","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sevanthi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Swarna","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Annayya","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Bharathi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"G","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Pushpa","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Raviraja","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sadananda","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Gopala","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"G","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"hemagiri","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"jayanthi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Karunakara","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"shivdas","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Shrimathi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"yuvaraj","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Yuvathikala","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Guruva","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Annaji","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Damodara","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Jagadish","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Malathi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Varija","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Vasantha","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Ramappa","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Anusuya","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Bojaraj","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Girraji","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Goverdan","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Hemavathi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Jayaprakash","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Jhanavi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Kasturi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Preema","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Somanatha","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Yadavi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sooru","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Bhuvanendra","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"Bhavya","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Brijesh","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Jithesh","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Yajnesh","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Chitra","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Prashant","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Shuba","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Hiriyanna","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Kalavathi","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Dhanaraj","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Geetha","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Jaya","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Prethiviraj","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Santosh","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Suma","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Lalitha","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Jyothi","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Shoba","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Leela","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Sujith","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Vineet","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sharada","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Ashish","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Devanand","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Mohini","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Shalinia","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Tharanatha","Gender":"male","Husband":null,"Wife":null,"children":[{"name":"B","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Minu","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Reeka","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Roopa","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Vanaja","Gender":"female","Husband":null,"Wife":null,"children":[{"name":"Devaki","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Giri","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Indira","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Rohit","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Sarala","Gender":"female","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""},{"name":"Shesha","Gender":"male","Husband":null,"Wife":null,"children":[],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""}],"Photo":"","Occupation":"","Accomplishments":null,"FacebookRef":"","TwitterRef":""};

// Set the dimensions and margins of the diagram
var margin = {top: 20, right: 90, bottom: 30, left: 90},
    width = 960 - margin.left - margin.right,
    height = 500 - margin.top - margin.bottom;

// append the svg object to the body of the page
// appends a 'group' element to 'svg'
// moves the 'group' element to the top left margin
var svg = d3.select("body").append("svg")
    .attr("width", width + margin.right + margin.left)
    .attr("height", height + margin.top + margin.bottom)
  .append("g")
    .attr("transform", "translate("
          + margin.left + "," + margin.top + ")");

var i = 0,
    duration = 750,
    root;

// declares a tree layout and assigns the size
var treemap = d3.tree().size([height, width]);

// Assigns parent, children, height, depth
root = d3.hierarchy(treeData, function(d) { return d.children; });
root.x0 = height / 2;
root.y0 = 0;

// Collapse after the second level
root.children.forEach(collapse);

update(root);

// Collapse the node and all it's children
function collapse(d) {
  if(d.children) {
    d._children = d.children
    d._children.forEach(collapse)
    d.children = null
  }
}

function update(source) {

  // Assigns the x and y position for the nodes
  var treeData = treemap(root);

  // Compute the new tree layout.
  var nodes = treeData.descendants(),
      links = treeData.descendants().slice(1);

  // Normalize for fixed-depth.
  nodes.forEach(function(d){ d.y = d.depth * 180});

  // ****************** Nodes section ***************************

  // Update the nodes...
  var node = svg.selectAll('g.node')
      .data(nodes, function(d) {return d.id || (d.id = ++i); });

  // Enter any new modes at the parent's previous position.
  var nodeEnter = node.enter().append('g')
      .attr('class', 'node')
      .attr("transform", function(d) {
        return "translate(" + source.y0 + "," + source.x0 + ")";
    })
    .on('click', click);

  // Add Circle for the nodes
  nodeEnter.append('circle')
      .attr('class', 'node')
      .attr('r', 1e-6)
      .style("fill", function(d) {
          return d._children ? "lightsteelblue" : "#fff";
      });

  // Add labels for the nodes
  nodeEnter.append('text')
      .attr("dy", ".35em")
      .attr("x", function(d) {
          return d.children || d._children ? -13 : 13;
      })
      .attr("text-anchor", function(d) {
          return d.children || d._children ? "end" : "start";
      })
      .text(function(d) { return d.data.name; });

  // UPDATE
  var nodeUpdate = nodeEnter.merge(node);

  // Transition to the proper position for the node
  nodeUpdate.transition()
    .duration(duration)
    .attr("transform", function(d) {
        return "translate(" + d.y + "," + d.x + ")";
     });

  // Update the node attributes and style
  nodeUpdate.select('circle.node')
    .attr('r', 10)
    .style("fill", function(d) {
        return d._children ? "lightsteelblue" : "#fff";
    })
    .attr('cursor', 'pointer');


  // Remove any exiting nodes
  var nodeExit = node.exit().transition()
      .duration(duration)
      .attr("transform", function(d) {
          return "translate(" + source.y + "," + source.x + ")";
      })
      .remove();

  // On exit reduce the node circles size to 0
  nodeExit.select('circle')
    .attr('r', 1e-6);

  // On exit reduce the opacity of text labels
  nodeExit.select('text')
    .style('fill-opacity', 1e-6);

  // ****************** links section ***************************

  // Update the links...
  var link = svg.selectAll('path.link')
      .data(links, function(d) { return d.id; });

  // Enter any new links at the parent's previous position.
  var linkEnter = link.enter().insert('path', "g")
      .attr("class", "link")
      .attr('d', function(d){
        var o = {x: source.x0, y: source.y0}
        return diagonal(o, o)
      });

  // UPDATE
  var linkUpdate = linkEnter.merge(link);

  // Transition back to the parent element position
  linkUpdate.transition()
      .duration(duration)
      .attr('d', function(d){ return diagonal(d, d.parent) });

  // Remove any exiting links
  var linkExit = link.exit().transition()
      .duration(duration)
      .attr('d', function(d) {
        var o = {x: source.x, y: source.y}
        return diagonal(o, o)
      })
      .remove();

  // Store the old positions for transition.
  nodes.forEach(function(d){
    d.x0 = d.x;
    d.y0 = d.y;
  });

  // Creates a curved (diagonal) path from parent to the child nodes
  function diagonal(s, d) {

    path = `M ${s.y} ${s.x}
            C ${(s.y + d.y) / 2} ${s.x},
              ${(s.y + d.y) / 2} ${d.x},
              ${d.y} ${d.x}`

    return path
  }

  // Toggle children on click.
  function click(d) {
    if (d.children) {
        d._children = d.children;
        d.children = null;
      } else {
        d.children = d._children;
        d._children = null;
      }
    update(d);
  }
}

</script>
</body>