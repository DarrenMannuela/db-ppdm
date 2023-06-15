package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfLanding(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_landing")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_landing

	for rows.Next() {
		var sf_landing dto.Sf_landing
		if err := rows.Scan(&sf_landing.Sf_subtype, &sf_landing.Support_facility_id, &sf_landing.Activation_freq_desc, &sf_landing.Activation_tone_desc, &sf_landing.Active_ind, &sf_landing.Airpspace_desc, &sf_landing.Altitude, &sf_landing.Altitude_ouom, &sf_landing.Approach_direction, &sf_landing.Area_id, &sf_landing.Area_size, &sf_landing.Area_size_ouom, &sf_landing.Area_type, &sf_landing.Communication_freq, &sf_landing.Communication_freq_desc, &sf_landing.Depart_direction, &sf_landing.Effective_date, &sf_landing.Expiry_date, &sf_landing.Fuel_avail_desc, &sf_landing.Landing_facility_code, &sf_landing.Landing_type, &sf_landing.Lighting_avail_ind, &sf_landing.Lighting_cycle_desc, &sf_landing.Lighting_desc, &sf_landing.Max_allow_mass, &sf_landing.Max_allow_mass_desc, &sf_landing.Max_allow_mass_ouom, &sf_landing.Perm_obstacle_desc, &sf_landing.Ppdm_guid, &sf_landing.Radio_call_name, &sf_landing.Radio_channel, &sf_landing.Remark, &sf_landing.Source, &sf_landing.Special_procedure_desc, &sf_landing.Surface_desc, &sf_landing.Surface_type, &sf_landing.Weather_info_desc, &sf_landing.Windcone_ind, &sf_landing.Row_changed_by, &sf_landing.Row_changed_date, &sf_landing.Row_created_by, &sf_landing.Row_created_date, &sf_landing.Row_effective_date, &sf_landing.Row_expiry_date, &sf_landing.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_landing to result
		result = append(result, sf_landing)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfLanding(c *fiber.Ctx) error {
	var sf_landing dto.Sf_landing

	setDefaults(&sf_landing)

	if err := json.Unmarshal(c.Body(), &sf_landing); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_landing values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45)")
	if err != nil {
		return err
	}
	sf_landing.Row_created_date = formatDate(sf_landing.Row_created_date)
	sf_landing.Row_changed_date = formatDate(sf_landing.Row_changed_date)
	sf_landing.Effective_date = formatDateString(sf_landing.Effective_date)
	sf_landing.Expiry_date = formatDateString(sf_landing.Expiry_date)
	sf_landing.Row_effective_date = formatDateString(sf_landing.Row_effective_date)
	sf_landing.Row_expiry_date = formatDateString(sf_landing.Row_expiry_date)






	rows, err := stmt.Exec(sf_landing.Sf_subtype, sf_landing.Support_facility_id, sf_landing.Activation_freq_desc, sf_landing.Activation_tone_desc, sf_landing.Active_ind, sf_landing.Airpspace_desc, sf_landing.Altitude, sf_landing.Altitude_ouom, sf_landing.Approach_direction, sf_landing.Area_id, sf_landing.Area_size, sf_landing.Area_size_ouom, sf_landing.Area_type, sf_landing.Communication_freq, sf_landing.Communication_freq_desc, sf_landing.Depart_direction, sf_landing.Effective_date, sf_landing.Expiry_date, sf_landing.Fuel_avail_desc, sf_landing.Landing_facility_code, sf_landing.Landing_type, sf_landing.Lighting_avail_ind, sf_landing.Lighting_cycle_desc, sf_landing.Lighting_desc, sf_landing.Max_allow_mass, sf_landing.Max_allow_mass_desc, sf_landing.Max_allow_mass_ouom, sf_landing.Perm_obstacle_desc, sf_landing.Ppdm_guid, sf_landing.Radio_call_name, sf_landing.Radio_channel, sf_landing.Remark, sf_landing.Source, sf_landing.Special_procedure_desc, sf_landing.Surface_desc, sf_landing.Surface_type, sf_landing.Weather_info_desc, sf_landing.Windcone_ind, sf_landing.Row_changed_by, sf_landing.Row_changed_date, sf_landing.Row_created_by, sf_landing.Row_created_date, sf_landing.Row_effective_date, sf_landing.Row_expiry_date, sf_landing.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfLanding(c *fiber.Ctx) error {
	var sf_landing dto.Sf_landing

	setDefaults(&sf_landing)

	if err := json.Unmarshal(c.Body(), &sf_landing); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_landing.Sf_subtype = id

    if sf_landing.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_landing set support_facility_id = :1, activation_freq_desc = :2, activation_tone_desc = :3, active_ind = :4, airpspace_desc = :5, altitude = :6, altitude_ouom = :7, approach_direction = :8, area_id = :9, area_size = :10, area_size_ouom = :11, area_type = :12, communication_freq = :13, communication_freq_desc = :14, depart_direction = :15, effective_date = :16, expiry_date = :17, fuel_avail_desc = :18, landing_facility_code = :19, landing_type = :20, lighting_avail_ind = :21, lighting_cycle_desc = :22, lighting_desc = :23, max_allow_mass = :24, max_allow_mass_desc = :25, max_allow_mass_ouom = :26, perm_obstacle_desc = :27, ppdm_guid = :28, radio_call_name = :29, radio_channel = :30, remark = :31, source = :32, special_procedure_desc = :33, surface_desc = :34, surface_type = :35, weather_info_desc = :36, windcone_ind = :37, row_changed_by = :38, row_changed_date = :39, row_created_by = :40, row_effective_date = :41, row_expiry_date = :42, row_quality = :43 where sf_subtype = :45")
	if err != nil {
		return err
	}

	sf_landing.Row_changed_date = formatDate(sf_landing.Row_changed_date)
	sf_landing.Effective_date = formatDateString(sf_landing.Effective_date)
	sf_landing.Expiry_date = formatDateString(sf_landing.Expiry_date)
	sf_landing.Row_effective_date = formatDateString(sf_landing.Row_effective_date)
	sf_landing.Row_expiry_date = formatDateString(sf_landing.Row_expiry_date)






	rows, err := stmt.Exec(sf_landing.Support_facility_id, sf_landing.Activation_freq_desc, sf_landing.Activation_tone_desc, sf_landing.Active_ind, sf_landing.Airpspace_desc, sf_landing.Altitude, sf_landing.Altitude_ouom, sf_landing.Approach_direction, sf_landing.Area_id, sf_landing.Area_size, sf_landing.Area_size_ouom, sf_landing.Area_type, sf_landing.Communication_freq, sf_landing.Communication_freq_desc, sf_landing.Depart_direction, sf_landing.Effective_date, sf_landing.Expiry_date, sf_landing.Fuel_avail_desc, sf_landing.Landing_facility_code, sf_landing.Landing_type, sf_landing.Lighting_avail_ind, sf_landing.Lighting_cycle_desc, sf_landing.Lighting_desc, sf_landing.Max_allow_mass, sf_landing.Max_allow_mass_desc, sf_landing.Max_allow_mass_ouom, sf_landing.Perm_obstacle_desc, sf_landing.Ppdm_guid, sf_landing.Radio_call_name, sf_landing.Radio_channel, sf_landing.Remark, sf_landing.Source, sf_landing.Special_procedure_desc, sf_landing.Surface_desc, sf_landing.Surface_type, sf_landing.Weather_info_desc, sf_landing.Windcone_ind, sf_landing.Row_changed_by, sf_landing.Row_changed_date, sf_landing.Row_created_by, sf_landing.Row_effective_date, sf_landing.Row_expiry_date, sf_landing.Row_quality, sf_landing.Sf_subtype)
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

func PatchSfLanding(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_landing set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where sf_subtype = :id"

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

func DeleteSfLanding(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_landing dto.Sf_landing
	sf_landing.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_landing where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_landing.Sf_subtype)
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


