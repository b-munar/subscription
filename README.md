# Instalacion

1. Clonar repositorio

```bash
git clone https://github.com/b-munar/subscription.git
```

2. 

```bash
docker compose build
docker compose up
```


El servicio de subscription.

## 1. Get list subscription

Retorna subscription.

<table>
<tr>
<td> Método </td>
<td> GET </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6450/subscription/list</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>N/A</td>
</tr>
<tr>
<td> Cuerpo </td>
<td>N/A</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 200 </td>
<td>En caso de exito</td>
<td>

```json
{
  "subscription": {
    "subscriptions": [
      {
        "plan": "basico",
        "price": 0
      },
      {
        "plan": "intermedio",
        "price": 19
      },
      {
        "plan": "premium",
        "price": 39
      }
    ],
    "activate": false
  }
}
```
</td>
</tr>
</tbody>
</table>

## 2. insert subscription

Retorna subscription.

<table>
<tr>
<td> Método </td>
<td> Post </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6450/subscription</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>"Authorization": "Bearer eyJ0eXAiOiJKV1QiL..."/td>
</tr>
<tr>
<td> Cuerpo </td>
<td>

```json
    {
      "plan": "premium",
      "price": 39
}
```
</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 201 </td>
<td>En caso de exito</td>
<td>

```json
{
  "subscription": {
    "subscription": {
      "id": "9a771fd8-ca52-4b80-a125-4f0eab9c6ec3",
      "price": 39,
      "frequency": 0,
      "plan": "premium",
      "created_at": "2024-04-13T00:30:17.59361276Z",
      "updated_at": "2024-04-13T00:30:17.59361276Z"
    },
    "activate": true
  }
}
```
</td>
</tr>
</tbody>
</table>



## 3. Get subscription

Retorna subscription.

<table>
<tr>
<td> Método </td>
<td> Get </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6450/subscription</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>"Authorization": "Bearer eyJ0eXAiOiJKV1QiL..."/td>
</tr>
<tr>
<td> Cuerpo </td>
<td>

NA

</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 200 </td>
<td>En caso de exito</td>
<td>

```json
{
  "subscription": {
    "subscription": {
      "id": "9a771fd8-ca52-4b80-a125-4f0eab9c6ec3",
      "price": 39,
      "frequency": 0,
      "plan": "premium",
      "created_at": "2024-04-13T00:30:17.59361276Z",
      "updated_at": "2024-04-13T00:30:17.59361276Z"
    },
    "activate": true
  }
}
```

or 

```json
{
  "subscription": {
    "subscriptions": [
      {
        "plan": "basico",
        "price": 0
      },
      {
        "plan": "intermedio",
        "price": 19
      },
      {
        "plan": "premium",
        "price": 39
      }
    ],
    "activate": false
  }
}
```

</td>
</tr>
</tbody>
</table>

