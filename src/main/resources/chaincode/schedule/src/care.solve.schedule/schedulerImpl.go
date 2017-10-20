package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"encoding/json"
	"errors"
)

type SchedulerImpl struct {
}

func (s SchedulerImpl) constructScheduleId(scheduleId string) string {
	return "schedule:" + scheduleId
}

func (s SchedulerImpl) Get(stub shim.ChaincodeStubInterface, scheduleId string) (*Schedule, error) {
	scheduleKey := s.constructScheduleId(scheduleId)
	scheduleBytes, err := stub.GetState(scheduleKey)

	if err != nil {
		return nil, err
	}

	var schedule Schedule
	if scheduleBytes == nil {
		return nil, errors.New(fmt.Sprintf("Schedule with id '%v' not found", scheduleId))
	} else {
		json.Unmarshal(scheduleBytes, &schedule)
		logger.Infof("Retrieve schedule: %v", schedule)
	}

	return &schedule, nil
}

func (s SchedulerImpl) Apply(stub shim.ChaincodeStubInterface, scheduleId string, scheduleRecord ScheduleRecord) error {
	scheduleKey := s.constructScheduleId(scheduleId)
	scheduleBytes, err := stub.GetState(scheduleKey)

	var schedule Schedule
	if scheduleBytes == nil {
		logger.Infof("Empty schedule for doctor %v. Creating new...", scheduleId)
		schedule = Schedule{scheduleKey, scheduleId, make([]*ScheduleRecord, 0)}
	} else {
		logger.Infof("Found schedule for doctor %v", scheduleId)
		json.Unmarshal(scheduleBytes, &schedule)
	}

	schedule.Records = append(schedule.Records, &scheduleRecord)

	jsonSchedule, err := json.Marshal(schedule)
	if err != nil {
		return err
	}

	err = stub.PutState(scheduleKey, jsonSchedule)
	if err != nil {
		return err
	}

	return nil
}