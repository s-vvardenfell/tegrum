#pragma once

#include <iostream>
#include <fstream>
#include <string>
#include <sstream>

#include <chrono>
#include <iomanip>
#include <sstream>

#include <vector>
#include <algorithm>
#include <map>

#include <nlohmann/json.hpp>
#include <tinyxml2.h>

#include "../../gumbo-query/headers/Document.h"
#include "../../gumbo-query/headers/Selector.h"
#include "../../gumbo-query/headers/Node.h"

using json = nlohmann::json;

using std::cerr;
using std::cout;
using std::endl;

std::string getTimestamp();
tinyxml2::XMLElement* getElementByName(tinyxml2::XMLDocument& doc, std::string const& elemt_value);
const std::string loadDocument(const std::string_view filename);
time_t getUnixTimeFromStr(const std::string& date_time, const std::string& format);
std::vector<tinyxml2::XMLNode*> getNodesByName(tinyxml2::XMLNode* root, std::string_view name);
std::string getDateTime(const std::string_view fmt);
bool saveDocument(const std::string& filename, const std::string& data);
void processProxie(const std::string& file);
std::vector<std::string> split(std::string& s, char ch);