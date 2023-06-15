package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetHseIncidentWeather(c *fiber.Ctx) error {
	rows, err := db.Query("select * from hse_incident_weather")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Hse_incident_weather

	for rows.Next() {
		var hse_incident_weather dto.Hse_incident_weather
		if err := rows.Scan(&hse_incident_weather.Incident_id, &hse_incident_weather.Weather_obs_no, &hse_incident_weather.Active_ind, &hse_incident_weather.Effective_date, &hse_incident_weather.Expiry_date, &hse_incident_weather.Heave, &hse_incident_weather.Pitch, &hse_incident_weather.Ppdm_guid, &hse_incident_weather.Recorded_time, &hse_incident_weather.Recorded_timezone, &hse_incident_weather.Remark, &hse_incident_weather.Road_condition, &hse_incident_weather.Roll, &hse_incident_weather.Source, &hse_incident_weather.Swell_height, &hse_incident_weather.Temperature, &hse_incident_weather.Temperature_ouom, &hse_incident_weather.Water_condition, &hse_incident_weather.Wave_height, &hse_incident_weather.Weather_condition, &hse_incident_weather.Wind_direction, &hse_incident_weather.Wind_strength, &hse_incident_weather.Row_changed_by, &hse_incident_weather.Row_changed_date, &hse_incident_weather.Row_created_by, &hse_incident_weather.Row_created_date, &hse_incident_weather.Row_effective_date, &hse_incident_weather.Row_expiry_date, &hse_incident_weather.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append hse_incident_weather to result
		result = append(result, hse_incident_weather)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetHseIncidentWeather(c *fiber.Ctx) error {
	var hse_incident_weather dto.Hse_incident_weather

	setDefaults(&hse_incident_weather)

	if err := json.Unmarshal(c.Body(), &hse_incident_weather); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into hse_incident_weather values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	hse_incident_weather.Row_created_date = formatDate(hse_incident_weather.Row_created_date)
	hse_incident_weather.Row_changed_date = formatDate(hse_incident_weather.Row_changed_date)
	hse_incident_weather.Effective_date = formatDateString(hse_incident_weather.Effective_date)
	hse_incident_weather.Expiry_date = formatDateString(hse_incident_weather.Expiry_date)
	hse_incident_weather.Recorded_time = formatDateString(hse_incident_weather.Recorded_time)
	hse_incident_weather.Row_effective_date = formatDateString(hse_incident_weather.Row_effective_date)
	hse_incident_weather.Row_expiry_date = formatDateString(hse_incident_weather.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_weather.Incident_id, hse_incident_weather.Weather_obs_no, hse_incident_weather.Active_ind, hse_incident_weather.Effective_date, hse_incident_weather.Expiry_date, hse_incident_weather.Heave, hse_incident_weather.Pitch, hse_incident_weather.Ppdm_guid, hse_incident_weather.Recorded_time, hse_incident_weather.Recorded_timezone, hse_incident_weather.Remark, hse_incident_weather.Road_condition, hse_incident_weather.Roll, hse_incident_weather.Source, hse_incident_weather.Swell_height, hse_incident_weather.Temperature, hse_incident_weather.Temperature_ouom, hse_incident_weather.Water_condition, hse_incident_weather.Wave_height, hse_incident_weather.Weather_condition, hse_incident_weather.Wind_direction, hse_incident_weather.Wind_strength, hse_incident_weather.Row_changed_by, hse_incident_weather.Row_changed_date, hse_incident_weather.Row_created_by, hse_incident_weather.Row_created_date, hse_incident_weather.Row_effective_date, hse_incident_weather.Row_expiry_date, hse_incident_weather.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateHseIncidentWeather(c *fiber.Ctx) error {
	var hse_incident_weather dto.Hse_incident_weather

	setDefaults(&hse_incident_weather)

	if err := json.Unmarshal(c.Body(), &hse_incident_weather); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	hse_incident_weather.Incident_id = id

    if hse_incident_weather.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update hse_incident_weather set weather_obs_no = :1, active_ind = :2, effective_date = :3, expiry_date = :4, heave = :5, pitch = :6, ppdm_guid = :7, recorded_time = :8, recorded_timezone = :9, remark = :10, road_condition = :11, roll = :12, source = :13, swell_height = :14, temperature = :15, temperature_ouom = :16, water_condition = :17, wave_height = :18, weather_condition = :19, wind_direction = :20, wind_strength = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where incident_id = :29")
	if err != nil {
		return err
	}

	hse_incident_weather.Row_changed_date = formatDate(hse_incident_weather.Row_changed_date)
	hse_incident_weather.Effective_date = formatDateString(hse_incident_weather.Effective_date)
	hse_incident_weather.Expiry_date = formatDateString(hse_incident_weather.Expiry_date)
	hse_incident_weather.Recorded_time = formatDateString(hse_incident_weather.Recorded_time)
	hse_incident_weather.Row_effective_date = formatDateString(hse_incident_weather.Row_effective_date)
	hse_incident_weather.Row_expiry_date = formatDateString(hse_incident_weather.Row_expiry_date)






	rows, err := stmt.Exec(hse_incident_weather.Weather_obs_no, hse_incident_weather.Active_ind, hse_incident_weather.Effective_date, hse_incident_weather.Expiry_date, hse_incident_weather.Heave, hse_incident_weather.Pitch, hse_incident_weather.Ppdm_guid, hse_incident_weather.Recorded_time, hse_incident_weather.Recorded_timezone, hse_incident_weather.Remark, hse_incident_weather.Road_condition, hse_incident_weather.Roll, hse_incident_weather.Source, hse_incident_weather.Swell_height, hse_incident_weather.Temperature, hse_incident_weather.Temperature_ouom, hse_incident_weather.Water_condition, hse_incident_weather.Wave_height, hse_incident_weather.Weather_condition, hse_incident_weather.Wind_direction, hse_incident_weather.Wind_strength, hse_incident_weather.Row_changed_by, hse_incident_weather.Row_changed_date, hse_incident_weather.Row_created_by, hse_incident_weather.Row_effective_date, hse_incident_weather.Row_expiry_date, hse_incident_weather.Row_quality, hse_incident_weather.Incident_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchHseIncidentWeather(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update hse_incident_weather set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "recorded_time" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where incident_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteHseIncidentWeather(c *fiber.Ctx) error {
	id := c.Params("id")
	var hse_incident_weather dto.Hse_incident_weather
	hse_incident_weather.Incident_id = id

	stmt, err := db.Prepare("delete from hse_incident_weather where incident_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(hse_incident_weather.Incident_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


