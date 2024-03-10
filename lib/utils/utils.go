package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)


func AnimateText(str string, animation []string, ms time.Duration, loop int) {
		maxReplace := len(str)
		loops := func() int {
			if loop < 1 {
				return 1
			}
			return loop
		}()

		for l := 0; l < loops; l++ {
			for i := 0; i < len(animation); i++ {
				if (len(animation[i]) > maxReplace) {
					maxReplace += len(animation[i])
				}
				if i > len(animation) - 1 {
					i = 0
				}

				fmt.Printf("\r  %s %s" + strings.Repeat(" ", maxReplace + 2) +"\r", str, animation[i])
				time.Sleep(ms * time.Millisecond)
			}
		}
	}
	func LoadingText(str string) {
			var	chars = []string{".   ", "..  ", "... ", "....", " ...", "  ..", "   .", "              "}
			AnimateText(str, chars, 65, 2)
		}

func CheckNilErr(e error) {
		if e != nil {
			log.Fatal("Error message:", e)
			return
		}
	}

func DisplayName(m *discordgo.Member) string {
		if m.Nick != "" {
			return m.Nick
		}
		return m.User.Username
	}

func FormatBool(b bool) string {
	    if b {
	        return "true"
	    }
	    return "false"
	}

func IsURL(str string) bool {
	    u, err := url.Parse(str)
	    return err == nil && u.Scheme != "" && u.Host != ""
	}

func PrettyPrint(data interface{}, ident string) {
	    var p []byte
	    p, err := json.MarshalIndent(data, "  ", "\t")
	    if err != nil {
	        fmt.Println(err)
	        return
	    }
	    fmt.Printf("  %s = ", ident)
	    fmt.Printf("%s \n\n", p)
	}

func StartupAnim(appID string, perms string) {
		LoadingText("Logging in")
		fmt.Println("\r  Bot Running.  Press CTRL-C to exit." + strings.Repeat(" ", 50))
		fmt.Printf("  https://discord.com/api/oauth2/authorize?client_id=%s&permissions=%s&scope=bot\n", appID, perms)
	}

func TypeOf(v interface{}) string {
	    return fmt.Sprintf("%T", v)
	}