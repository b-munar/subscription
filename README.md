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
<td> <strong>localhost:6450/subscription</strong> </td>
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
  "subscription": [
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
  ]
}
```
</td>
</tr>
</tbody>
</table>
