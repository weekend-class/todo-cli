//
//  main.swift
//  Todo
//
//  Created by Rakhasa on 06/02/23.
//

import Foundation

var command: [String] = [add, ls, rm, uncheck, check]
var data: [DataStruct] = []

var input : String = ""

func getCheckIcon (isChecked: Bool) -> String {
    return isChecked ? "✅" : ""
}

print("ToDo CLI")

repeat {
    print("Input some command : ", terminator: "")
    input = readLine() ?? ""

    var commandList : [String] = []

    for element in command {
        if input.lowercased().contains(element) {
            commandList.append(element)
        }
    }

    if commandList.isEmpty {
        print("command list (add, ls, rm, check, uncheck, quit)")
        continue
    }

    switch commandList.first {

    case add:
        let split = input.components(separatedBy: "\(add) ")
        if split.count > 1 {
            data.append(DataStruct(name : split[1],isChecked: false ))
        }

    case ls:
        for (index, element) in data.enumerated() {
            print("\(index + 1). \(element.name) \(getCheckIcon(isChecked: element.isChecked)) ")
        }

    case rm:
        let split = input.components(separatedBy: "\(rm) ")
        if split.count > 1 {
            let index : Int? = Int(split[1])

            if index == nil {
                continue
            }

            if data.count <= (index! - 1) {
                print("no list at number \(String(index!))")
                continue
            }

            data.remove(at: index! - 1)
        }

    case check:
        let split = input.components(separatedBy: "\(check) ")
        if split.count > 1 {
            let index : Int? = Int(split[1])

            if index == nil {
                continue
            }

            if data.count <= (index! - 1) {
                print("no list at number \(String(index!))")
                continue
            }

            data[index! - 1].isChecked = true
        }

    case uncheck:
        let split = input.components(separatedBy: "\(uncheck) ")
        if split.count > 1 {
            let index : Int? = Int(split[1])

            if index == nil {
                continue
            }

            if data.count <= (index! - 1) {
                print("no list at number \(String(index!))")
                continue
            }

            data[index! - 1].isChecked = false
        }

    default:
        exit(0)
    }

} while (input != "quit")

