# Cliente MQTT implementado para el Trabajo Práctico Anual de DDS

Véase el enunciado [TP Anual DDS](https://docs.google.com/document/d/13niiEppxrm8LjyrxmH5Pskrc7VVuPKWSFRi3WvhsXns/edit?tab=t.0)

## Descripción

La implementación permite simular los mensajes que enviaría los sensores de una heladera. Los tópics
se dividen por heladera, y para diferenciar el sensor se define una estructura de mensaje como las
siguientes:

```bash
TEMPERATURA <numero>

FRAUDE

SOLICITUD_APERTURA <código de tarjeta>
```

## Comandos de consola

`quit` para salir y terminar la ejecución.

### Envío de Temperatura

```bash
temp [NUMBER] [TOPIC]
```

El argumento `NUMBER` es opcional, en caso de no especificar, simplemente se envía una temperatura
random entre 2 y -2.

### Alerta de Fraude

```bash
fraud [TOPIC]
```

### Solicitud de Apertura

```bash
open-for [CODE] [TOPIC]
```

El argumento `CODE` es obligatorio, es el código de la tarjeta la cual está realizando la solicitud.

## Configuración y Variables de Entorno

Las opciones disponibles para el `.env` se encuentran en el archivo `.env.schema`.

Un ejemplo de la config podría ser la siguiente:

```conf
HELADERA_TOPICS="
    utn-dds-g22/heladeras/heladera-plaza-de-mayo,
    utn-dds-g22/heladeras/heladera-ferro,
    utn-dds-g22/heladeras/heladera-atlanta,
    utn-dds-g22/heladeras/heladera-utn-lugano,
    utn-dds-g22/heladeras/heladera-obelisco,
    utn-dds-g22/heladeras/heladera-caminito-de-la-boca,
    utn-dds-g22/heladeras/heladera-plaza-italia,
    utn-dds-g22/heladeras/heladera-barrio-chino,
    utn-dds-g22/heladeras/heladera-facultad-de-derecho,
    utn-dds-g22/heladeras/heladera-hospital-piniero,
    utn-dds-g22/heladeras/heladera-plaza-serrano,
    utn-dds-g22/heladeras/heladera-abasto-shopping,
    utn-dds-g22/heladeras/heladera-guerrin,
    utn-dds-g22/heladeras/heladera-linea-d,
    utn-dds-g22/heladeras/heladera-utn-medrano"
PUBLISH_TEMPERATURE=true
PUBLISH_TEMPERATURE_INTERVAL=2
```

### Publicación de Temperatura automática

Seteando `PUBLISH_TEMPERATURE` en `true`, el cliente mqtt publicará de forma automática temperaturas
random a los topics definidos en `HELADERA_TOPICS`.

La variable `PUBLISH_TEMPERATURE_INTERVAL` indica el retardo en segundos entre publicaciones de
temperatura del mismo topic; el intervalo no puede ser 0, se seteará en 1 si se especifica cualquier
número menor a éste. En caso de no definirlo, tendrá un valor de 5 por defecto.
