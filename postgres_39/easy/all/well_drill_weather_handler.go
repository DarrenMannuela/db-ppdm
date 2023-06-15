package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillWeather(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_weather")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_weather

	for rows.Next() {
		var well_drill_weather dto.Well_drill_weather
		if err := rows.Scan(&well_drill_weather.Uwi, &well_drill_weather.Period_obs_no, &well_drill_weather.Weather_obs_no, &well_drill_weather.Active_ind, &well_drill_weather.Effective_date, &well_drill_weather.Expiry_date, &well_drill_weather.Heave, &well_drill_weather.Pitch, &well_drill_weather.Ppdm_guid, &well_drill_weather.Recorded_time, &well_drill_weather.Recorded_timezone, &well_drill_weather.Remark, &well_drill_weather.Road_condition, &well_drill_weather.Roll, &well_drill_weather.Source, &well_drill_weather.Swell_height, &well_drill_weather.Temperature, &well_drill_weather.Temperature_ouom, &well_drill_weather.Water_condition, &well_drill_weather.Wave_height, &well_drill_weather.Weather_condition, &well_drill_weather.Wind_direction, &well_drill_weather.Wind_strength, &well_drill_weather.Row_changed_by, &well_drill_weather.Row_changed_date, &well_drill_weather.Row_created_by, &well_drill_weather.Row_created_date, &well_drill_weather.Row_effective_date, &well_drill_weather.Row_expiry_date, &well_drill_weather.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_weather to result
		result = append(result, well_drill_weather)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillWeather(c *fiber.Ctx) error {
	var well_drill_weather dto.Well_drill_weather

	setDefaults(&well_drill_weather)

	if err := json.Unmarshal(c.Body(), &well_drill_weather); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_weather values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	well_drill_weather.Row_created_date = formatDate(well_drill_weather.Row_created_date)
	well_drill_weather.Row_changed_date = formatDate(well_drill_weather.Row_changed_date)
	well_drill_weather.Effective_date = formatDateString(well_drill_weather.Effective_date)
	well_drill_weather.Expiry_date = formatDateString(well_drill_weather.Expiry_date)
	well_drill_weather.Recorded_time = formatDateString(well_drill_weather.Recorded_time)
	well_drill_weather.Row_effective_date = formatDateString(well_drill_weather.Row_effective_date)
	well_drill_weather.Row_expiry_date = formatDateString(well_drill_weather.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_weather.Uwi, well_drill_weather.Period_obs_no, well_drill_weather.Weather_obs_no, well_drill_weather.Active_ind, well_drill_weather.Effective_date, well_drill_weather.Expiry_date, well_drill_weather.Heave, well_drill_weather.Pitch, well_drill_weather.Ppdm_guid, well_drill_weather.Recorded_time, well_drill_weather.Recorded_timezone, well_drill_weather.Remark, well_drill_weather.Road_condition, well_drill_weather.Roll, well_drill_weather.Source, well_drill_weather.Swell_height, well_drill_weather.Temperature, well_drill_weather.Temperature_ouom, well_drill_weather.Water_condition, well_drill_weather.Wave_height, well_drill_weather.Weather_condition, well_drill_weather.Wind_direction, well_drill_weather.Wind_strength, well_drill_weather.Row_changed_by, well_drill_weather.Row_changed_date, well_drill_weather.Row_created_by, well_drill_weather.Row_created_date, well_drill_weather.Row_effective_date, well_drill_weather.Row_expiry_date, well_drill_weather.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillWeather(c *fiber.Ctx) error {
	var well_drill_weather dto.Well_drill_weather

	setDefaults(&well_drill_weather)

	if err := json.Unmarshal(c.Body(), &well_drill_weather); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_weather.Uwi = id

    if well_drill_weather.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_weather set period_obs_no = :1, weather_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, heave = :6, pitch = :7, ppdm_guid = :8, recorded_time = :9, recorded_timezone = :10, remark = :11, road_condition = :12, roll = :13, source = :14, swell_height = :15, temperature = :16, temperature_ouom = :17, water_condition = :18, wave_height = :19, weather_condition = :20, wind_direction = :21, wind_strength = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where uwi = :30")
	if err != nil {
		return err
	}

	well_drill_weather.Row_changed_date = formatDate(well_drill_weather.Row_changed_date)
	well_drill_weather.Effective_date = formatDateString(well_drill_weather.Effective_date)
	well_drill_weather.Expiry_date = formatDateString(well_drill_weather.Expiry_date)
	well_drill_weather.Recorded_time = formatDateString(well_drill_weather.Recorded_time)
	well_drill_weather.Row_effective_date = formatDateString(well_drill_weather.Row_effective_date)
	well_drill_weather.Row_expiry_date = formatDateString(well_drill_weather.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_weather.Period_obs_no, well_drill_weather.Weather_obs_no, well_drill_weather.Active_ind, well_drill_weather.Effective_date, well_drill_weather.Expiry_date, well_drill_weather.Heave, well_drill_weather.Pitch, well_drill_weather.Ppdm_guid, well_drill_weather.Recorded_time, well_drill_weather.Recorded_timezone, well_drill_weather.Remark, well_drill_weather.Road_condition, well_drill_weather.Roll, well_drill_weather.Source, well_drill_weather.Swell_height, well_drill_weather.Temperature, well_drill_weather.Temperature_ouom, well_drill_weather.Water_condition, well_drill_weather.Wave_height, well_drill_weather.Weather_condition, well_drill_weather.Wind_direction, well_drill_weather.Wind_strength, well_drill_weather.Row_changed_by, well_drill_weather.Row_changed_date, well_drill_weather.Row_created_by, well_drill_weather.Row_effective_date, well_drill_weather.Row_expiry_date, well_drill_weather.Row_quality, well_drill_weather.Uwi)
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

func PatchWellDrillWeather(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_weather set "
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
	query += " where uwi = :id"

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

func DeleteWellDrillWeather(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_weather dto.Well_drill_weather
	well_drill_weather.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_weather where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_weather.Uwi)
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


