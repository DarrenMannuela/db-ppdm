package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacility(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility

	for rows.Next() {
		var facility dto.Facility
		if err := rows.Scan(&facility.Facility_id, &facility.Facility_type, &facility.Abandoned_date, &facility.Active_date, &facility.Active_ind, &facility.Catalogue_equip_id, &facility.Constructed_date, &facility.Coord_acquisition_id, &facility.Coord_system_id, &facility.Current_operator, &facility.Depth, &facility.Depth_ouom, &facility.Description, &facility.Effective_date, &facility.Elevation, &facility.Elevation_ouom, &facility.Expiry_date, &facility.Facility_diameter, &facility.Facility_diameter_ouom, &facility.Facility_function, &facility.Facility_long_name, &facility.Facility_no, &facility.Facility_short_name, &facility.H2s_ind, &facility.Inactive_date, &facility.Last_injection_date, &facility.Last_production_date, &facility.Last_reported_date, &facility.Latitude, &facility.Local_coord_system_id, &facility.Longitude, &facility.Manufactured_by, &facility.On_injection_date, &facility.On_production_date, &facility.Pipeline_material, &facility.Pipeline_type, &facility.Pipe_cover_type, &facility.Plot_name, &facility.Plot_symbol, &facility.Ppdm_guid, &facility.Primary_field_id, &facility.Remark, &facility.Source, &facility.Unit_operated_ind, &facility.Row_changed_by, &facility.Row_changed_date, &facility.Row_created_by, &facility.Row_created_date, &facility.Row_effective_date, &facility.Row_expiry_date, &facility.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility to result
		result = append(result, facility)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacility(c *fiber.Ctx) error {
	var facility dto.Facility

	setDefaults(&facility)

	if err := json.Unmarshal(c.Body(), &facility); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51)")
	if err != nil {
		return err
	}
	facility.Row_created_date = formatDate(facility.Row_created_date)
	facility.Row_changed_date = formatDate(facility.Row_changed_date)
	facility.Abandoned_date = formatDateString(facility.Abandoned_date)
	facility.Active_date = formatDateString(facility.Active_date)
	facility.Constructed_date = formatDateString(facility.Constructed_date)
	facility.Effective_date = formatDateString(facility.Effective_date)
	facility.Expiry_date = formatDateString(facility.Expiry_date)
	facility.Inactive_date = formatDateString(facility.Inactive_date)
	facility.Last_injection_date = formatDateString(facility.Last_injection_date)
	facility.Last_production_date = formatDateString(facility.Last_production_date)
	facility.Last_reported_date = formatDateString(facility.Last_reported_date)
	facility.On_injection_date = formatDateString(facility.On_injection_date)
	facility.On_production_date = formatDateString(facility.On_production_date)
	facility.Row_effective_date = formatDateString(facility.Row_effective_date)
	facility.Row_expiry_date = formatDateString(facility.Row_expiry_date)






	rows, err := stmt.Exec(facility.Facility_id, facility.Facility_type, facility.Abandoned_date, facility.Active_date, facility.Active_ind, facility.Catalogue_equip_id, facility.Constructed_date, facility.Coord_acquisition_id, facility.Coord_system_id, facility.Current_operator, facility.Depth, facility.Depth_ouom, facility.Description, facility.Effective_date, facility.Elevation, facility.Elevation_ouom, facility.Expiry_date, facility.Facility_diameter, facility.Facility_diameter_ouom, facility.Facility_function, facility.Facility_long_name, facility.Facility_no, facility.Facility_short_name, facility.H2s_ind, facility.Inactive_date, facility.Last_injection_date, facility.Last_production_date, facility.Last_reported_date, facility.Latitude, facility.Local_coord_system_id, facility.Longitude, facility.Manufactured_by, facility.On_injection_date, facility.On_production_date, facility.Pipeline_material, facility.Pipeline_type, facility.Pipe_cover_type, facility.Plot_name, facility.Plot_symbol, facility.Ppdm_guid, facility.Primary_field_id, facility.Remark, facility.Source, facility.Unit_operated_ind, facility.Row_changed_by, facility.Row_changed_date, facility.Row_created_by, facility.Row_created_date, facility.Row_effective_date, facility.Row_expiry_date, facility.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacility(c *fiber.Ctx) error {
	var facility dto.Facility

	setDefaults(&facility)

	if err := json.Unmarshal(c.Body(), &facility); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility.Facility_id = id

    if facility.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility set facility_type = :1, abandoned_date = :2, active_date = :3, active_ind = :4, catalogue_equip_id = :5, constructed_date = :6, coord_acquisition_id = :7, coord_system_id = :8, current_operator = :9, depth = :10, depth_ouom = :11, description = :12, effective_date = :13, elevation = :14, elevation_ouom = :15, expiry_date = :16, facility_diameter = :17, facility_diameter_ouom = :18, facility_function = :19, facility_long_name = :20, facility_no = :21, facility_short_name = :22, h2s_ind = :23, inactive_date = :24, last_injection_date = :25, last_production_date = :26, last_reported_date = :27, latitude = :28, local_coord_system_id = :29, longitude = :30, manufactured_by = :31, on_injection_date = :32, on_production_date = :33, pipeline_material = :34, pipeline_type = :35, pipe_cover_type = :36, plot_name = :37, plot_symbol = :38, ppdm_guid = :39, primary_field_id = :40, remark = :41, source = :42, unit_operated_ind = :43, row_changed_by = :44, row_changed_date = :45, row_created_by = :46, row_effective_date = :47, row_expiry_date = :48, row_quality = :49 where facility_id = :51")
	if err != nil {
		return err
	}

	facility.Row_changed_date = formatDate(facility.Row_changed_date)
	facility.Abandoned_date = formatDateString(facility.Abandoned_date)
	facility.Active_date = formatDateString(facility.Active_date)
	facility.Constructed_date = formatDateString(facility.Constructed_date)
	facility.Effective_date = formatDateString(facility.Effective_date)
	facility.Expiry_date = formatDateString(facility.Expiry_date)
	facility.Inactive_date = formatDateString(facility.Inactive_date)
	facility.Last_injection_date = formatDateString(facility.Last_injection_date)
	facility.Last_production_date = formatDateString(facility.Last_production_date)
	facility.Last_reported_date = formatDateString(facility.Last_reported_date)
	facility.On_injection_date = formatDateString(facility.On_injection_date)
	facility.On_production_date = formatDateString(facility.On_production_date)
	facility.Row_effective_date = formatDateString(facility.Row_effective_date)
	facility.Row_expiry_date = formatDateString(facility.Row_expiry_date)






	rows, err := stmt.Exec(facility.Facility_type, facility.Abandoned_date, facility.Active_date, facility.Active_ind, facility.Catalogue_equip_id, facility.Constructed_date, facility.Coord_acquisition_id, facility.Coord_system_id, facility.Current_operator, facility.Depth, facility.Depth_ouom, facility.Description, facility.Effective_date, facility.Elevation, facility.Elevation_ouom, facility.Expiry_date, facility.Facility_diameter, facility.Facility_diameter_ouom, facility.Facility_function, facility.Facility_long_name, facility.Facility_no, facility.Facility_short_name, facility.H2s_ind, facility.Inactive_date, facility.Last_injection_date, facility.Last_production_date, facility.Last_reported_date, facility.Latitude, facility.Local_coord_system_id, facility.Longitude, facility.Manufactured_by, facility.On_injection_date, facility.On_production_date, facility.Pipeline_material, facility.Pipeline_type, facility.Pipe_cover_type, facility.Plot_name, facility.Plot_symbol, facility.Ppdm_guid, facility.Primary_field_id, facility.Remark, facility.Source, facility.Unit_operated_ind, facility.Row_changed_by, facility.Row_changed_date, facility.Row_created_by, facility.Row_effective_date, facility.Row_expiry_date, facility.Row_quality, facility.Facility_id)
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

func PatchFacility(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility set "
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
		} else if key == "abandoned_date" || key == "active_date" || key == "constructed_date" || key == "effective_date" || key == "expiry_date" || key == "inactive_date" || key == "last_injection_date" || key == "last_production_date" || key == "last_reported_date" || key == "on_injection_date" || key == "on_production_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where facility_id = :id"

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

func DeleteFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility dto.Facility
	facility.Facility_id = id

	stmt, err := db.Prepare("delete from facility where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility.Facility_id)
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


