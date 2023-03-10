# 鈿機upcake Core
Cupcake Core es el empaquetado de software que brinda servicios de comunicaci贸n
y transferencia de datos entre los clientes que hacen uso de la **Suite de Cupcake**.

El core, en su primera versi贸n, implementa un servicio REST que permite almacenar y consultar 
informaci贸n de la base de datos de cupcake.

## 馃摑Caracter铆sticas
* Sistema de registros concurrente
* Conexi贸n con base de datos MySQL
* Servidor REST

## 馃搵 Implementaciones
El Core esta pensado para funcionar en entornos de peque帽as y medianas empresas,
provey茅ndolas de servicios escalables, por lo que cada una es libre de realizar una implementaci贸n local o distribuida seg煤n sean sus necesidades.

### 馃搷 Implementaci贸n local [En desarrollo]
<img src="docs/images/imple-1.png" alt="">
Esta implementaci贸n permite a la empresa concentrar el flujo de datos dentro de su propia
organizaci贸n, pudiendo o no, exponer los servicios de Cupcake Core para que otros usuarios fuera 
de su 谩rea de implementaci贸n puedan hacer uso de estos.

### 馃寪 Implementaci贸n distribuida [En desarrollo]
<img src="docs/images/imple-2.png" alt="">
En la implementaci贸n distribuida, Cupcake Core dota a la empresa con posibilidad de escalar su
flujo de datos en diferentes 谩reas de implementaci贸n.
Ya sea que se encuentren en diferentes zonas geogr谩ficas, cada una de las implementaciones 
podr谩 hacer uso de una copia de la base de datos de la infraestructura central, pero cada una 
podr谩 hacer uso de sus propios servicios.

Asimismo, Cupcake no se limita a una 煤nica instancia dentro de una infraestructura,
sino que adem谩s se puede dotar a la infraestructura con m煤ltiples instancias de 
servicios escuchando en diferentes puertos. La configuraci贸n de esta queda a criterio de 
cada empres

## 馃ЬLicencia
Cupcake Core se encuentra bajo licencia MIT, por lo que cualquiera puede 
hacer uso de este core para implementarlo en sus propios proyectos, respetando las propias directrices de la licencia.
